# Buscador de Archivos en Go

Un buscador de archivos para Linux desarrollado en Go. Recorre directorios de forma recursiva utilizando **goroutines** para realizar búsquedas paralelas, logrando tiempos de ejecución rapidísimos incluso en sistemas de archivos grandes.

## 🚀 Objetivo

El objetivo principal del proyecto fue practicar el uso de **goroutines** y comprender mejor cómo aprovechar el paralelismo en Go para crear herramientas eficientes desde la línea de comandos (CLI).

## 🔧 Características

- Búsqueda recursiva de archivos
- Concurrencia con goroutines para mayor velocidad
- Exclusión de rutas específicas
- Uso simple desde la terminal

## 🛠️ Instalación y uso

Cloná el repositorio:

```bash
git clone https://github.com/FedericoDeniard/buscador-archivos.git
cd buscador-archivos
```

Compilá e instalá el programa usando `make`:

```bash
make build
make install
```

Luego podés usar el comando `buscador` directamente desde cualquier lugar:

```bash
buscador -file=mi_archivo.txt -deep -exclude=/proc,/sys
```

Para desinstalarlo:

```bash
sudo make uninstall
```

### 🏷️ Flags disponibles

| Flag       | Descripción                                                      |
| ---------- | ---------------------------------------------------------------- |
| `-file`    | **(Obligatorio)** Archivo que se desea buscar                    |
| `-deep`    | Inicia la búsqueda desde el directorio raíz del sistema (`/`)    |
| `-exclude` | Lista de directorios a excluir, separados por comas sin espacios |
| `-help`    | Muestra el mensaje de ayuda                                      |

## 🤝 Contribuciones

¡Las contribuciones son bienvenidas! Si encontrás un bug o querés proponer mejoras, sentite libre de abrir un issue o un pull request.

---

Desarrollado con Go 🦫 por [Federico Deniard](https://github.com/FedericoDeniard)
