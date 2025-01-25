package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
)

type huffman_node struct{
  freq int;
  char_data byte;
  left_node *huffman_node;
  right_node *huffman_node;
}

func main(){
  fileNames := os.Args[1:]
  fmt.Printf("Reading your file....\n")
  for _,fileName := range fileNames{ 
    fileData,err := os.ReadFile(fileName)
    if err !=nil {
      fmt.Fprintf(os.Stderr,"[ERROR]:"+err.Error())
    }
    encode_into_huffman(string(fileData))
  }
}


func encode_into_huffman(fileData string){
  // "helloworld"
  // h->1 e->1 l->3 o->2 w->1 r->1 d->1
  resulting_map := map[byte]int {}
 // Create a frequency list
  for _,value := range fileData{
    if value != 0b1010{ // 01010 -> \n
    resulting_map[byte(value)]++;
    }
  }
  //Debugging
  
  /*for key_value,value := range resulting_map {
   huf_node := huffman_node { value , key_value , nil ,nil}
   fmt.Printf("Character:%v\t%v\n",huf_node.freq,string(huf_node.char_data));
  } */

  sorted_int := make([][2]interface{},0,len(resulting_map)) // interface 
 
  for k,v := range resulting_map{
    sorted_int = append(sorted_int, [2]interface{}{k,v})
  }
  sort.Slice(sorted_int,func(i,j int)bool{
    return sorted_int[i][1].(int) < sorted_int[j][1].(int)
  })

  keys := make([]byte, len(sorted_int))  // @ sorted keys
  
   for i, p := range sorted_int{
      keys[i] = p[0].(byte) // sets interface array's 0th element to be key
  }

  for _,key := range keys{
    fmt.Printf("Frequency:%v\t Character:%v\n",resulting_map[byte(key)],string(key)) 
  }

  // now that I have a huffman node for each value, first I will sort the map with respect to the frequency and create a tree
}
func merge_node(n1* huffman_node,n2* huffman_node) huffman_node{
  return huffman_node{n1.freq+n2.freq,0b0110,n1,n2}
}
