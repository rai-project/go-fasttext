package main

import (
  "os"
  "bufio"
  "log"
  "fmt"
  "strings"
  "strconv"
  "sort"
  "math"
  "gonum.org/v1/gonum/mat"
)

type Pair struct {
    Key   string
    Value float64
}
type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }

func main() {
  argc := len(os.Args)
  var dim int

  if argc <= 2{
    if os.Args[1] == "analogy"{
      println("usage: go run evaluate.py analogy path/to/your/model.txt path/to/your/testfile")
      println("   or: go run evaluate.py analogy path/to/your/model.txt word1 word2 word3")
    }
    if os.Args[1] == "nn"{
      println("usage: go run evaluate.py nn path/to/your/model.txt word")
    }
  } else {
    model := os.Args[2]
    file, err := os.Open(model)
    if err != nil {
      log.Fatal(err)
    }
    defer file.Close()

    m := make(map[string][]float64)
    m_dot := make(map[string]float64)

    var name_vector[]string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      line := scanner.Text()
      elements := strings.Fields(line)
      vector := elements[1:]
      dim = len(vector)
      name_vector = append(name_vector, elements[0])
      var vector_float[]float64
      sum := 0.0
      for i := 0; i < dim; i++ {
        number,err := strconv.ParseFloat(vector[i],64)
        if err != nil {
          log.Fatal(err)
        }
        vector_float = append(vector_float, number)
        sum += math.Pow(number,2)
      }
      sum = math.Sqrt(sum)
      // normalize word vectors
      for i := 0; i < dim; i++ {
        vector_float[i] = vector_float[i] / sum
      }
      m[elements[0]] = vector_float
    }
    if err := scanner.Err(); err != nil {
      log.Fatal(err)
    }

    if os.Args[1] == "analogy"{
      if argc == 6{
        word1 := mat.NewVecDense(dim, m[os.Args[3]])
        word2 := mat.NewVecDense(dim, m[os.Args[4]])
        word3 := mat.NewVecDense(dim, m[os.Args[5]])
        result_vector := mat.NewVecDense(dim, nil)
        result_vector.AddVec(word1, word3)
        result_vector.SubVec(result_vector, word2)
        for k := range m{
          target_word := mat.NewVecDense(dim, m[k])
          d := mat.Dot(result_vector, target_word)
          m_dot[k] = d
        }
        pl := make(PairList, len(m_dot))
        i := 0
        for k, v := range m_dot{
          if k == os.Args[3] || k == os.Args[4] || k == os.Args[5]{
            continue
          }
          pl[i] = Pair{k,v}
          i++
        }
        sort.Sort(sort.Reverse(pl))
        for i = 0; i < 10; i++{
          fmt.Println(pl[i])
        }
      } else{
        var testfile[14]string
        testfile[0] = "analogy_data/capital-common-countries.txt"
        testfile[1] = "analogy_data/capital-world.txt"
        testfile[2] = "analogy_data/currency.txt"
        testfile[3] = "analogy_data/city-in-state.txt"
        testfile[4] = "analogy_data/family.txt"
        testfile[5] = "analogy_data/gram1-adjective-to-adverb.txt"
        testfile[6] = "analogy_data/gram2-opposite.txt"
        testfile[7] = "analogy_data/gram3-comparative.txt"
        testfile[8] = "analogy_data/gram4-superlative.txt"
        testfile[9] = "analogy_data/gram5-present-participle.txt"
        testfile[10] = "analogy_data/gram6-nationality-adjective.txt"
        testfile[11] = "analogy_data/gram7-past-tense.txt"
        testfile[12] = "analogy_data/gram8-plural.txt"
        testfile[13] = "analogy_data/gram9-plural-verbs.txt"
        for i := 0; i < 14; i++{
          file, err := os.Open(testfile[i])
          if err != nil {
            log.Fatal(err)
          }
          defer file.Close()
          scanner := bufio.NewScanner(file)
          numer := 0
          denom := 0
          for scanner.Scan() {
            denom ++
            line := scanner.Text()
            elements := strings.Fields(line)
            word1 := mat.NewVecDense(dim, m[elements[0]])
            word2 := mat.NewVecDense(dim, m[elements[1]])
            word3 := mat.NewVecDense(dim, m[elements[3]])
            word4 := elements[2]
            result_vector := mat.NewVecDense(dim, nil)
            result_vector.AddVec(word1, word3)
            result_vector.SubVec(result_vector, word2)
            for k := range m {
              target_word := mat.NewVecDense(dim, m[k])
              d := mat.Dot(result_vector, target_word)
              m_dot[k] = d
            }
            pl := make(PairList, len(m_dot))
            i := 0
            for k, v := range m_dot{
              if k == elements[0] || k == elements[1] || k == elements[3]{
                continue
              }
              pl[i] = Pair{k,v}
              i++
            }
            sort.Sort(sort.Reverse(pl))
            if pl[0].Key == word4{
              numer ++
            }
          }
          if err := scanner.Err(); err != nil {
            log.Fatal(err)
          }
          fmt.Printf("%s: %f\n", testfile[i], float64(numer)/float64(denom))
        }
      }
    }
    if os.Args[1] == "nn"{
    }
  }
}

// analogy_data/capital-common-countries.txt: 0.470356
// analogy_data/capital-world.txt: 0.164898
// analogy_data/currency.txt: 0.195150
// analogy_data/city-in-state.txt: 0.100932
// analogy_data/family.txt: 0.422925
// analogy_data/gram1-adjective-to-adverb.txt: 0.017137
// analogy_data/gram2-opposite.txt: 0.038177
// analogy_data/gram3-comparative.txt: 0.335586
// analogy_data/gram4-superlative.txt: 0.138146
// analogy_data/gram5-present-participle.txt: 0.228220
// analogy_data/gram6-nationality-adjective.txt: 0.772358
// analogy_data/gram7-past-tense.txt: 0.123077
// analogy_data/gram8-plural.txt: 0.289790
// analogy_data/gram9-plural-verbs.txt: 0.226437
