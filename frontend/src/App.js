import { useRef } from 'react';
import Editor from '@monaco-editor/react';
import './App.css';

function App() {
  const dir = "http://localhost:5000/";
  var nombreArchivo = "";
  const editorRef = useRef(null);

  function handleEditorDidMount(editor, monaco) {
    editorRef.current = editor;
  }

  function abrirArchivo(e) {
    var archivo = e.target.files[0];
    nombreArchivo = archivo.name
    const fileReader = new FileReader();
    fileReader.readAsText(archivo);
    fileReader.addEventListener('load', (e) => {
      editorRef.current.getModel().setValue(fileReader.result);
    });
  }

  function guardar() {
    if (nombreArchivo != "") {
      var archivo = new Blob([editorRef.current.getValue()], { type: 'text/plain' });
      if (window.navigator.msSaveOrOpenBlob)
        window.navigator.msSaveOrOpenBlob(archivo, nombreArchivo);
      else {
        let a = document.createElement("a"),
        url = URL.createObjectURL(archivo);
        a.href = url;
        a.download = nombreArchivo;
        document.body.appendChild(a);
        a.click();
        setTimeout(function () {
          document.body.removeChild(a);
          window.URL.revokeObjectURL(url);
        }, 0);
      }
    }else {
      alert("No se ha cargado ningun archivo.\n\n Se creará uno nuevo.");
      guardarComo();
    }
  }

  function guardarComo() {
    let nuevoNombre = window.prompt("Ingrese el nombre del archivo:", "Nuevo Archivo.swift");
    if (nuevoNombre) {
      nombreArchivo = nuevoNombre;
      let archivo = new Blob([editorRef.current.getValue()], { type: 'text/plain' });
      let a = document.createElement("a"),
      url = URL.createObjectURL(archivo);
      a.href = url;
      a.download = nombreArchivo;
      document.body.appendChild(a);
      a.click();
      setTimeout(function () {
        document.body.removeChild(a);
        window.URL.revokeObjectURL(url);
      }, 0);
    } else {
      alert("No se ingreso ningun nombre");
    }
  }

  function ejecutar() {
    var txt = { 'contenido': editorRef.current.getValue()}
    console.log(txt)
    fetch(`${dir}ejecutar`, {
      method: 'POST',
      body: JSON.stringify(txt)
    })
      .then(res => res.json())
      .catch(err => {
        console.error('Error:', err)
        alert("Error")
      })
      .then(response => {
        console.log("respuesta: ",response);
        document.querySelector('#consola').value = response.Salida;
      })
  }

  function errores() {
    fetch(`${dir}errores`, {
      method: 'GET',
    })
      .then(res => res.json())
      .catch(err => {
        console.error('Error:', err)
        alert("Error")
      })
      .then(response => {
        console.log(response);
        let listaErrores = []
        if (response.Salida == null) {
          alert("No hay errores")
          return
        }
        listaErrores = response.Salida
        var datos = `
        <!DOCTYPE html>
        <html lang="es">
        <head>
            <meta charset="UTF-8">
            <meta http-equiv="X-UA-Compatible" content="IE=edge">
            <title> REPORTE ERRORES </title>
            <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta3/dist/css/bootstrap.min.css" rel="stylesheet"
                integrity="sha384-eOJMYsd53ii+scO/bJGFsiCZc+5NDVN2yr8+0RDqr0Ql0h+rP48ckxlpbzKgwra6" crossorigin="anonymous">
        </head>
        <body >        
            <div style="position: absolute; width: 98%; left: 1%; top: 20px;">
            <h1 style="margin-bottom: 20px;"> <center> REPORTE DE ERRORES </center></h1>
                <table class="table table-ligth table-striped table-hover table-bordered border-dark">
                    <thead class="table table-dark">
                        <tr>
                            <th scope="col">TIPO</th>
                            <th scope="col">DESCRIPCION</th>
                            <th scope="col">LINEA</th>
                            <th scope="col">COLUMNA</th>
                        </tr>
                    </thead>
                    <tbody>`
          listaErrores.forEach(error => {
          datos += ` <tr>
            <td scope="col"> ${error.Tipo} </td>
            <td scope="col"> ${error.Mensaje}</td>
            <td scope="col"> ${error.Linea}</td>
            <td scope="col"> ${error.Columna}</td>
          </tr>`
              });
          datos += `                    </tbody>
          </table>
          </div>
          </body>
          </html>`

        var win = window.open('', '', 'height=700,width=750');
        win.document.write(datos);
        win.document.close();
      })

  }

  function ts() {
    fetch(`${dir}ts`, {
      method: 'GET',
    })
      .then(res => res.json())
      .catch(err => {
        console.error('Error:', err)
        alert("Error")
      })
      .then(response => {
        console.log(response);
        let tabla = response.Salida;
        console.log(tabla.Variables)
        const keys_variables = Object.keys(tabla.Variables);
        const keys_funciones = Object.keys(tabla.Funciones);
        console.log(tabla.Variables instanceof Map);
          let anterior = tabla.Anterior
          let nombre = tabla.Nombre
          var datos = `
          <!DOCTYPE html>
          <html lang="es">
          <head>
              <meta charset="UTF-8">
              <meta http-equiv="X-UA-Compatible" content="IE=edge">
              <title> Tabla de Simbolos -  ${tabla.nombre} </title>
              <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta3/dist/css/bootstrap.min.css" rel="stylesheet"
                  integrity="sha384-eOJMYsd53ii+scO/bJGFsiCZc+5NDVN2yr8+0RDqr0Ql0h+rP48ckxlpbzKgwra6" crossorigin="anonymous">
          </head>
          <body >        
              <div style="position: absolute; width: 98%; left: 1%; top: 20px;">
              <h1 style="margin-bottom: 20px;"> <center> TABLA DE SIMBOLOS </center></h1>
                  <table class="table table-ligth table-striped table-hover table-bordered border-dark">
                      <thead class="table table-dark">
                          <tr>
                              <th scope="col">AMBITO</th>
                              <th scope="col">ID</th>
                              <th scope="col">TIPO SIMBOLO</th>
                              <th scope="col">TIPO DATO</th>
                              
                              <th scope="col">LINEA</th>
                              <th scope="col">COLUMNA</th>

                          </tr>
                      </thead>
                      <tbody>
  `       
           for (const clave of keys_variables) {
            const valor = tabla.Variables[clave];
            datos += ` <tr>
            <td scope="col"> ${nombre} </td>
            <td scope="col"> ${clave} </td>
            <td scope="col"> Variable </td>
            <td scope="col"> ${obtenerTipo(valor.Tipo)}</td>
            <td scope="col"> ${valor.Linea}</td>
            <td scope="col"> ${valor.Columna}</td>
            </tr>`
          }

          if (nombre === "Global") {
            for (const clave of keys_funciones) {
              const valor = tabla.Funciones[clave];
              datos += ` <tr>
              <td scope="col"> ${nombre} </td>
              <td scope="col"> ${clave} </td>
              <td scope="col"> Función </td>
              <td scope="col"> ${obtenerTipo(valor.TipoReturn)}</td>
              <td scope="col"> ${valor.Linea}</td>
              <td scope="col"> ${valor.Columna}</td>
              </tr>`
            }
          }


/*           while (anterior !== undefined) {
            try {
              nombre = anterior.Nombre;
              simbolos = anterior.Variables
              metodos = anterior.Funciones
            } catch (error) {
              console.log(error)
            }
            simbolos.forEach(simbolo => {
              datos += ` <tr>
                <td scope="col"> ${nombre} </td>
                <td scope="col"> ${simbolo.id} </td>
                <td scope="col"> ${simbolo.tipoVar}</td>
                <td scope="col"> ${simbolo.tipo}</td>
                <td scope="col"> ${JSON.stringify(simbolo.valor)}</td>
              </tr>`
            });
            if (nombre === "Global") {
              metodos.forEach(metodo => {
                let tipo = ""
                if (metodo.tipoReturn === "VOID") {
                  tipo = "MÉTODO"
                } else {
                  tipo = "FUNCIÓN"
                }
                datos += ` <tr>
                  <td scope="col"> ${nombre} </td>
                  <td scope="col"> ${metodo.id} </td>
                  <td scope="col"> ${tipo} </td>
                  <td scope="col"> ${metodo.tipoReturn}</td>
                  <td scope="col"> </td>
                </tr>`
              });
            }


            anterior = anterior.anterior;
          } */
          datos += `                    </tbody>
                  </table>
                  </div>
                  </body>
                  </html>`

          var win = window.open('', '', 'height=700,width=850');
          win.document.write(datos);
          win.document.close();
        
      })

  }

  function obtenerTipo (tipo) {
    switch (tipo) {
      case 0:
        return "String"
      case 1:
        return "Bool"
      case 2:
        return "Character"
      case 3:
        return "Int"
      case 4:
        return "Float"
      case 6:
        return "Void"
    }  
  }

  return (
    <>
    <nav className="navbar navbar-expand-lg navbar-dark bg-dark">
        <div className="container">
            <a className="navbar-brand" href=""> <h4>T-Swift</h4> </a>
            <div className="collapse navbar-collapse" id="navbarNav">
                <ul className="navbar-nav ms-auto">
                <li className="nav-item">
                    <button className="btn btn-dark nav-link" type="button" data-bs-toggle="modal" data-bs-target="#cargar"> Abrir Archivo</button>
                </li>
                <li>
                  <button className="btn btn-dark nav-link" onClick={guardar}> Guardar </button>
                </li>
                <li>
                  <button className="btn btn-dark nav-link" onClick={guardarComo}> Guardar Como </button>
                </li>
                <li className="nav-item dropdown">
                  <a className="nav-link dropdown-toggle " href="" id="Reportes" role="button" data-bs-toggle="dropdown" aria-expanded="false">
                    Reportes
                  </a>
                  <ul className="dropdown-menu">
                    <li><a className="dropdown-item" role="button" onClick={errores}>Errores</a></li>
                    <li><a className="dropdown-item" role="button" onClick={ts}>Tabla de Símbolos</a></li>                    
                  </ul>
                </li>
                <li>
                    <button className="btn btn-primary" type="button" onClick={ejecutar}> Ejecutar </button>
                </li>
                </ul>
            </div>
        </div>
    </nav>
    <div className="modal fade" id="cargar" data-bs-backdrop="static" data-bs-keyboard="false" tabIndex="-1" aria-labelledby="staticBackdropLabel" aria-hidden="true">
        <div className="modal-dialog modal-dialog-centered">
            <div className="modal-content">
                <div className="modal-header">
                    <h5 className="modal-title" id="staticBackdropLabel">Cargar Archivo</h5>
                    <button type="button" className="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <div className="modal-body">
                    <input type="file" className="form-control" aria-label="Upload" accept=".swift" onChange={abrirArchivo} required />
                </div>
                <div className="modal-footer">
                    <button type="button" className="btn btn-primary" data-bs-dismiss="modal">Ok</button>
                </div>
            </div>
        </div>
    </div>

    <div className="modal fade" id="reportes" data-bs-backdrop="static" data-bs-keyboard="false" tabIndex="-1" aria-labelledby="staticBackdropLabel" aria-hidden="true">
      <div className="modal-dialog modal-fullscreen">
          <div className="modal-content">
              <div className="modal-header">
                  <h5 className="modal-title" id="staticBackdropLabel">Reportes</h5>
                  <button type="button" className="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
              </div>
              <div className="modal-body">
                <div style={{textAlign: 'center'}} id="Imagenes"></div>
              </div>
          </div>
      </div>
    </div>

    <form id="formEditor">
        <label htmlFor="Entrada" className="form-label titulos">Editor de Código</label>
        <Editor
          height="78vh"
          defaultLanguage="swift"
          defaultValue='print("Cristian Daniel Pereira Tezagüic - 202010893")'
          onMount={handleEditorDidMount}
          theme="vs-dark" />
    </form>
    <form id="formConsola">
        <label htmlFor="Salida" className="form-label titulos" > Salida </label>
        <textarea className=" form-control" id="consola" wrap="off" readOnly> </textarea>
        <div style={{marginBottom: '15px'}}></div>
    </form>
    </>
  );
}

export default App;
