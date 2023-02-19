# funtemps
Her er min versjon av Installasjon og Testing via funtemps.  

## Hvordan kjøre denne kommandoen
```code go
funtemps -C 20 -out F
```
Vil konvertere fra Celsius til Fahrenheit. Alle temperatur enheter som en kan bruke er Celsius, Fahrenheit og Kelvin.

## Kjøre kommando som "funtemps" og ikke "go run main.go"
For å kjøre funtemps må en putte funtemps.exe i PATH environment lokalt i Go i bin-mappen. Etter du har gjort det kan du kjøre denne terminal-kommandoen i terminal. 

## Begrunnelser
I denne koden valgte jeg å bruke switch and case uttrykk mot if og else uttrykk. Disse uttrykkene er brukt til å evaluere verdien av -out variabelen og kan dermed finne ut hvilken konverteringsfunksjon den vil kalle based på flaggene som er definert. Switch-uttrykk som vi lærte tidligere i semesteret gjør det enklere å lese koden samtidig som å gjøre det enklere å håndtere flere case-uttrykk med forskjellig output. I dette tilfellet er switch-uttrykk brukt til å finne ut hvilken type grad brukeren har lyst å konvertere til. Case-uttrykk er brukt til å sjekke om input flagg var gått igjennom og kaller den rette konverteringsfunksjonen. 

Når det kom til testing var det bruk av "go run main.go" kommandoen som jeg brukte for å teste output av forskjellige temperaturenheter. Jeg ble veldig usikker på conv_test.go siden jeg hadde implementert alle formler i conv.go, men jeg brukte conv_test.go til å kjøre som kommando via å konvertere den til en .exe fil. 

## Utfordringer
Min største utfordring var at jeg slet med å forstå hva som måtte gjøres og hva som måtte implementeres. Jeg følte det var veldig lite svar å få når det kom til selve oppgaven, jeg følte rett og slett at det var veldig uklart om hvor spesifikk vi måtte være i hva vi skulle gjøre, så jeg gjorde det som var mest naturlig for meg. 

En annen utfordring var jo selvsagt implementering av go.mod og hvordan jeg skulle implementere andre go-filer i main.go. Heldigvis var LA-ene til god hjelp her og jeg fant fort ut at jeg måtte bruke conv_test.go til å implementere conv-pakken. Dette føler jeg har vært veldig uklart men heldigvis gikk vi igjennom dette på seminaret så jeg fant ut av det til slutt.

Fun facts ble etterhvert ikke et krav og derfor valgte jeg det bort for å spare tid slik at jeg kunne fokusere på å gjøre selve switch og case uttrykkene bedre.

## Konklusjon
Jeg føler en stor mestringsfølelse av å utføre denne oppgaven. Det var en stor vegg å komme over, men jeg følte jeg klarte det til slutt.
