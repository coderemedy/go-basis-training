# Go Basis Training Übungen

## Übung 1 - Hello World
Lege ein neues go Projekt an und erstelle ein neues go Programm. Gebe mit Hilfe des ```fmt``` packages eine Ausgabe auf der Kommandozeile aus.


<details><summary>Hinweise</summary>

Hilfreiche Befehle:
- Package Initialisierung: ```go mod init <package name>```
- Build: ```go build .```
- Run: ```go run .```

</details>

## Übung 2 - Variablen, Funktionen und Errors
Lege in dieser Übung zwei Funktionen an.
- Die erste Funktion soll zwei Zahlen dividieren und das Ergebnis zurückgeben. Wenn der Divisor 0 ist, soll ein Fehler zurückgegeben werden. Nutze hierfür das ```fmt``` Package
- Die zweite Funktion soll einen als Pointer übergebenen String um einen weiteren String ergänzen.

Rufe beide Funktionen aus der main Funktion auf und lass die Ergebnisse aus der Konsole ausgeben. Ein eventuell auftretener Fehler soll beim Aufruf abgefangen werden und auf der Konsole geloggt werden.

<details><summary>Hinweise</summary>

- Nutze das ```fmt``` Package um einen Fehler zu erzeugen

</details>

## Übung 3 - Slices
Erstelle ein Slice eines beliebigen einfachen Types (int, string, ...) und fülle es mit beliebigen Werten. Probiere unterschiedliche Möglichkeiten der Slice Initialisierung und Befüllung aus. 
Gebe die Werte in einem Range Loop aus. 

## Übung 4 - Objektorientierung & Interfaces
Füge einem struct (entweder dem person struct, oder erstelle ein eigenes) ein Feld vom typ uint hinzu und implementiere das `GoStringer` interface für dieses struct. Gib dir anschließend das struct mit `fmt.Printf("%#v")` in der Kommandozeile aus.
Die ausgabe sollte den dezimalwert des uint feldes enthalten, nicht den hexadezimal wert!

<details><summary>Hinweise</summary>

Hilfreiche Referenzen & Befehle:
- https://pkg.go.dev/fmt#hdr-Printing.
- https://pkg.go.dev/fmt?utm_source=gopls#GoStringer.
- Build: ```go build .```.
- Run: ```go run .```.

</details>

## Übung 5 - Tests
Erstelle für eine Methode aus der Übung 2 einen Unit Test. Nutze If-Anweisungen und das fmt Package zur Validierung der Rückgabewerte.

<details><summary>Hinweise</summary>

- Methodensignatur: ```func TestXXX(t *testing.T) {}```
- Ausführen mit: ```go test ./...```

</details>

## Übung 6 - Concurrency
Erstelle eine Funktion die einen Wert (string, int, oder ähnliches) in einen Channel schreibt und und rufe sie als goroutine auf. Konsumiere das Resultat durch ein `select` statement und gebe den empfangenen Wert aus.

## Bonus Übung - Dependencies
Ändere den bereits erstellten Test aus der Übung 5 ab und nutze Assertions aus dem Testify Modul ([https://github.com/stretchr/testify](https://github.com/stretchr/testify)). Füge zusätzlich die hinzugefügten Dependencies deinem Modul lokal hinzu.

<details><summary>Hinweise</summary>

Hilfreiche Befehle:
- Package hinzufügen: ```go get -t "<package name>"```
- Modul Sync: ```go mod tidy```
- Vendoring: ```go mod vendor```

</details>
