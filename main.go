package main

import "fmt"

func main() {
  fmt.Println(transcribeDNAtoRNA("TCGAGGTACTTCGA"))
}

func translateDNAtoRNA(dna string) string {
  var rna string = ""
  for i := 0; i < len(dna); i++ {
    if dna[i] == 'T' {
      rna = rna + "U"
    } else {
      rna = rna + string(dna[i])
    }
  }
  return rna
}

func transcribeDNAtoRNA(dna string) string {
  var rna string = ""
  for i := 0; i < len(dna); i++ {
    switch dna[i] {
      case 'A':
        rna += "U"
      case 'T':
        rna += "A"
      case 'C':
        rna += "G"
      case 'G':
        rna += "C"
    }
  }
  return rna
}