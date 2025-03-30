// Package main è il punto di ingresso per un'applicazione Go eseguibile
package main

// Import della libreria standard per l'input/output
import (
    "fmt"        // Formattazione di testo e I/O
    "time"       // Funzionalità relative al tempo
    "strings"    // Manipolazione delle stringhe
)

// La funzione main è il punto di partenza dell'esecuzione
func main() {
    // Dichiarazione di variabile con tipo esplicito
    var messaggio string = "Hello, World!"
    
    // Dichiarazione di variabile con inferenza di tipo
    ora := time.Now()
    
    // Utilizzo di una funzione della libreria standard
    messaggioMaiuscolo := strings.ToUpper(messaggio)
    
    // Stampa semplice
    fmt.Println(messaggio)
    
    // Stampa formattata
    fmt.Printf("In maiuscolo: %s\n", messaggioMaiuscolo)
    
    // Utilizzo di una struttura dati (slice)
    parole := []string{"Hello", "World", "in", "Go!"}
    
    // Iterazione con range
    for i, parola := range parole {
        fmt.Printf("Parola %d: %s\n", i, parola)
    }
    
    // Utilizzo di un if con dichiarazione inline
    if orarioFormattato := ora.Format("15:04:05"); ora.Hour() < 12 {
        fmt.Printf("Buongiorno! Sono le %s\n", orarioFormattato)
    } else {
        fmt.Printf("Buonasera! Sono le %s\n", orarioFormattato)
    }
}