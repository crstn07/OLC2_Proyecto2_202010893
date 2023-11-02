# MANUAL TECNICO

## Requerimientos
- Golang
- ANTLR4
- Sistema Operativo indiferente

## Código Servidor en GO
### Structs

#### Mensaje
Estructura utilizada para guardar el codigo ingresado en el editor. <br>
```go
type Mensaje struct {
	Contenido string
}
```

#### Response
Estructura utilizada para guardar la respuesta de la tabla de simbolos. <br>
```go
type Response struct {
	Salida Scope
}
```

#### ResponseErrores
Estructura utilizada para guardar la respuesta de los errores. <br>
```go
type ResponseErrores struct {
	Salida []Error_
}
```

#### Respuesta
Estructura utilizada para guardar la respuesta de la ejecución del codigo. <br>
```go
type Respuesta struct {
	Salida string
}
```

#### Scope
Estructura utilizada para guardar la tabla de simbolos. <br>
```go
type Scope struct {
	Nombre       string
	Anterior     *Scope
	Variables    map[string]Variable
	Funciones    map[string]Funcion
	Size         map[string]int
	etqSalida    string
	etqCiclo     string
}

```

#### Variable
Estructura utilizada para guardar las variables. <br>
```go
type Variable struct {
	ID        string
	Tipo      int
	Valor     interface{}
	Constante bool
  Posicion  int
	Linea     int
	Columna   int
}
```

#### Funcion
Estructura utilizada para guardar las funciones. <br>
```go
type Funcion struct {
	ID         string
	Parametros parser.IParametrosContext
	Bloque     parser.IBlockContext
	TipoReturn int
	Linea      int
	Columna    int
}
```

#### Parametro
Estructura utilizada para guardar los parámetros de las funciones. <br>
```go
type Parametro struct {
	ID_Externo string
	ID_Interno string
	Tipo       int
	Referencia bool
}
```

#### Error_
Estructura utilizada para guardar los errores. <br>
```go
type Error_ struct {
	Tipo    string
	Linea   string
	Columna string
	Mensaje string
}
```

#### Valor
Estructura utilizada para guardar los valores de las variables. <br>
```go
type Valor struct {
	Valor      interface{}
	Tipo       int
}
```

#### Visitor
```go
type Visitor struct {
	parser.SwiftGrammarVisitor
}

```

#### Lista de Errors
Almacena todos los errores que se vayan detectando durante la ejecución
```go
var listaErrores []Error_
```

### Funciones

#### Main
En el main se crea un nuevo servidor mux que es el que se encargara de manejar las peticiones. <br>
Se crean los handlers de cada ruta que se utilizara en la API. <br>
El handler de /ejecutar es el encargado de ejecutar el codigo ingresado en el editor. <br>
El handler de /ts es el encargado de mostrar la tabla de simbolos. <br>
El handler de /errores es el encargado de mostrar los errores durante la ejecución. <br>
Se imprime en consola el puerto en el que se esta ejecutando el servidor. <br>
Se crea un nuevo handler para el cors y se le pasa el handler de mux. <br>
Se ejecuta el servidor en el puerto 5000. <br>
```go
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ejecutar", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		listaErrores = listaErrores[:0]
		var codigo Mensaje
		body, _ := io.ReadAll(r.Body)
		json.Unmarshal(body, &codigo)
		input := antlr.NewInputStream(codigo.Contenido)
		lexer := parser.NewSwiftGrammarLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewSwiftGrammarParser(stream)
		p.BuildParseTrees = true
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
		tree := p.Prog()
		eval := Visitor{}
		ts := Scope{Variables: make(map[string]Variable), Nombre: "Global", Anterior: nil, Funciones: make(map[string]Funcion)}
		res := eval.Visit(tree, ts).Valor
		tablaSimbolos = ts
		w.WriteHeader(http.StatusOK)
		respuesta := Respuesta{
			Salida: fmt.Sprintf("%v", res),
		}
		jsonData, err := json.Marshal(respuesta)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		CreateCST(codigo.Contenido)
		w.Write(jsonData)
	})

	mux.HandleFunc("/ts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		respuesta := Response{
			Salida: tablaSimbolos,
		}
		jsonData, err := json.Marshal(respuesta)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(jsonData)
		w.Write(jsonData)

	})

	mux.HandleFunc("/errores", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		respuesta := ResponseErrores{
			Salida: listaErrores,
		}
		jsonData, err := json.Marshal(respuesta)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(jsonData)
		w.Write(jsonData)

	})

	fmt.Println("Servidor corriendo en el puerto 5000")
	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
```

#### Visit
El visit es el encargado de recorrer el arbol de derivación, este recibe como parametros el arbol de analisis sintactico y la tabla de simbolos. <br>
El arbol que se recibe es el generado por el parser de ANTLR. <br>
Se hace un switch para saber que tipo de contexto se esta visitando. <br>
Cada contexto tiene su propia función en el visitor, estas funciones devuelven un valor. <br>
El valor devuelto por el visitor es el valor de la variable, el cual puede ser cualquier tipo de dato. <br>
```go
func (v *Visitor) Visit(tree antlr.ParseTree, ts Scope) Valor {
	switch val := tree.(type) {
	case *parser.ProgContext:
		return v.VisitProg(val, ts)
	case *parser.BlockContext:
		return v.VisitBlock(val, ts)
	case *parser.InstrContext:
		return v.VisitInstr(val, ts)
	case *parser.DeclaracionContext:
		return v.VisitDeclaracion(val, ts)
	case *parser.DeclaracionTipoValorContext:
		return v.VisitDeclaracionTipoValor(val, ts)
	case *parser.DeclaracionTipoContext:
		return v.VisitDeclaracionTipo(val, ts)
	case *parser.DeclaracionValorContext:
		return v.VisitDeclaracionValor(val, ts)
	case *parser.AsignacionContext:
		return v.VisitAsignacion(val, ts)
	case *parser.Print_instrContext:
		return v.VisitPrint_instr(val, ts)
	case *parser.If_instrContext:
		return v.VisitIf_instr(val, ts)
	case *parser.IfContext:
		return v.VisitIf(val, ts)
	case *parser.IfElseContext:
		return v.VisitIfElse(val, ts)
	case *parser.ElseIfContext:
		return v.VisitElseIf(val, ts)
	case *parser.Switch_instrContext:
		return v.VisitSwitch(val, ts)
	case *parser.While_instrContext:
		return v.VisitWhile_instr(val, ts)
	case *parser.For_instrContext:
		return v.VisitFor(val, ts)
	case *parser.GuardContext:
		return v.VisitGuard(val, ts)
	case *parser.FuncionContext:
		return v.VisitFuncion(val, ts)
	case *parser.Llamada_funcContext:
		return v.VisitLlamada_func(val, ts)
	case *parser.LlamadaFuncExprContext:
		return v.VisitLlamadaFuncExpr(val, ts)
	case *parser.Dec_vectorContext:
		return v.VisitDeclaracion_vector(val, ts)
	case *parser.Copia_vectorContext:
		return v.VisitCopia_vector(val, ts)
	case *parser.Modificacion_vectorContext:
		return v.VisitModificacion_vector(val, ts)
	case *parser.AppendContext:
		return v.VisitAppend(val, ts)
	case *parser.RemoveLastContext:
		return v.VisitRemoveLast(val, ts)
	case *parser.RemoveAtContext:
		return v.VisitRemoveAt(val, ts)
	case *parser.UmenosExprContext:
		return v.VisitUMenosExpr(val, ts)
	case *parser.NotExprContext:
		return v.VisitNotExpr(val, ts)
	case *parser.OpExprContext:
		return v.VisitOpExpr(val, ts)
	case *parser.ParExprContext:
		return v.VisitParExpr(val, ts)
	case *parser.IdExprContext:
		return v.VisitIdExpr(val, ts)
	case *parser.StrExprContext:
		return v.VisitStrExpr(val, ts)
	case *parser.IntExprContext:
		return v.VisitIntExpr(val, ts)
	case *parser.FloatExprContext:
		return v.VisitFloatExpr(val, ts)
	case *parser.BoolExprContext:
		return v.VisitBoolExpr(val, ts)
	case *parser.CharExprContext:
		return v.VisitCharExpr(val, ts)
	case *parser.NilExprContext:
		return v.VisitNilExpr(val, ts)
	case *parser.IntCastExprContext:
		return v.VisitIntCastExpr(val, ts)
	case *parser.FloatCastExprContext:
		return v.VisitFloatCastExpr(val, ts)
	case *parser.StringCastExprContext:
		return v.VisitStringCastExpr(val, ts)
	case *parser.AccesoVectorContext:
		return v.VisitAccesoVector(val, ts)
	case *parser.IsEmptyContext:
		return v.VisitIsEmpty(val, ts)
	case *parser.CountContext:
		return v.VisitCount(val, ts)
	default:
		return Valor{fmt.Sprintf("Contexto No reconocido: %v", val), Error, false, false, false, nil, 0}
	}

}
```

## Código del Cliente en Javascript

### Funciones Usadas
- abrirArchivo()
- guardar()
- guardarComo()
- ejecutar()
- errores()
- ts()
- obtenerTipo()
```js
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
                              <th scope="col">VALOR</th>
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
            <td scope="col"> ${JSON.stringify(valor.Valor)}</td>
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
              <td scope="col"> </td>
              <td scope="col"> ${valor.Linea}</td>
              <td scope="col"> ${valor.Columna}</td>
              </tr>`
            }
          }

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
```