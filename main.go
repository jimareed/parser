package main

import "fmt"
import "regexp"
import "os"
import "bufio"


func main() {

    var inputRegex, outputFormat [1]string

    inputRegex[0] = "(?P<school>.+) (\\[)(?P<type>\\S+)(\\]) (?P<location>\\S+) (.+)(\\:) (?P<count>\\d+)"
    outputFormat[0] = "${school}|${type}|${location}|${count}"

    // open input file
    file, err := os.Open("/data/input.txt")
    if err != nil {
        fmt.Println("error opening file")
    } else {
        defer file.Close()
    }

    // create output file
    fileout, err := os.Create("/data/output.txt")
    if (err != nil) {
        fmt.Println("error creating file")
    } else {
        defer fileout.Close()
    }
    
    re := regexp.MustCompile(inputRegex[0])

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        if (re.MatchString(scanner.Text())) {
            fmt.Println()

            _, err := fileout.WriteString(re.ReplaceAllString(scanner.Text(), outputFormat[0]))

            if (err != nil) {
                fmt.Println("error writing to file")
            }            
            _, err = fileout.WriteString("\n")
        } else {
            //
        }
    }
    
    if err := scanner.Err(); err != nil {
        fmt.Println("error reading file")
    }
    
    fileout.Sync()
      
}


