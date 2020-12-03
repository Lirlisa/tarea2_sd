Integrantes:
Harold Caballero; 201773602-k
Nicolás Castillo; 201773561-9

Roles:
Máquina 1 (dist45): namenode, cliente
Máquina 2 (dist47): datanode
Máquina 3 (dist46): datanode
Máquina 4 (dist48): datanode

Para el correcto funcionamiento del sistema, namenode debe ser el primero en activarse, luego los datanodes y
finalmente el cliente, el cual puede ser activado las veces que se requieran.

Namenode:
El namenode se encuentra $HOME/tarea2_sd/Namenode, una vez ahí se puede ejecutar con el comando 'make run'.
El log queda almacenado en el archivo 'LOG.txt'
Se modificó ligeramente la estructura del log quedando:
<nombre libro> <Cantidad de chunks>
<nombre libro>_<número chunk> <ip maquina>

Datanode:
El datanode se encuentra en $HOME/tarea2_sd/, una vez ahí se puede ejecutar con el comando 'make run'.
Al recibir chunks, estos se almacenarán con el nombre '<nombre libro>_<número chunk>', sin extensión

Cliente:
El cliente se encuentra en $HOME/tarea2_sd/Cliente, una vez ahí se puede ejecutar con el comando 'make run'.
Para subir archivos, se debe entregar el nombre exacto del archivo pdf con su extensión, hay archivos ya cargados
para probar. Al elegir el modo descarga, se debe escribir el string exacto que se presenta por consola para hacer
la solicitud correctamente, luego el archivo descargado tendrá la forma 'Descarga<nombre del libro>.pdf'

Los archivos diponibles para prueba son:
- PeterPan.pdf
- Mujercitas.pdf
- Ivanhoe.pdf
- Dracula.pdf
- CuentodeNavidad.pdf

Sólo la parte distribuida fue implementada.
