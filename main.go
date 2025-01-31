package main

import (
	"fmt"
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
  if len(fileNames)<=0 {
    fmt.Fprintf(os.Stderr,"[ERROR]:No files found!\n");
    return
  }

  for _,fileName := range fileNames{ 
    fileData,err := os.ReadFile(fileName)
    if err !=nil {
      fmt.Fprintf(os.Stderr,"[ERROR]:"+err.Error())
    }
    resulting_arr := encode_into_huffman(string(fileData)) //now the thing returns the sorted array of huffman_nodes
    final_root_huffman_node := create_huff_node_tree(resulting_arr)
    fmt.Printf("[FINAL-TREE-ROOT]-%v\n",final_root_huffman_node)
    traverse_root(&final_root_huffman_node);
  }
}



func create_huff_node_tree(huffman_node_arr []huffman_node)huffman_node{
  // so, what needs to be done is that the huffman node array have to be converted to a tree like structure.
  // what i plan to do is, to loop through the array and for each new node that i get after merge the smallest two huffman_nodes
  // when pushing i just push and return the new sorted array and keep on doing that until the array is one element wide.
  // merge_node() function to merge two nodes - helper function
  // push_and_sort() function to push into the initial array and return the sorted array
  fmt.Printf("[NON-MERGED-TREE]-%v\n",huffman_node_arr)
  resulting_huff_node_arr := huffman_node_arr
  for ;len(resulting_huff_node_arr)!=1; {   
  resulting_huff_node_arr =  push_and_sort(merge_node(&resulting_huff_node_arr[0],&resulting_huff_node_arr[1]),resulting_huff_node_arr[2:])
  }
  return resulting_huff_node_arr[0] 
}

func encode_into_huffman(fileData string) []huffman_node{
  // "helloworld"
  // h->1 e->1 l->3 o->2 w->1 r->1 d->1
  var huffman_node_arr []huffman_node; 
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
      keys[i] = p[0].(byte) // type assertion,ie takes the value given that it is a key.
      // sets interface array's 0th element to be key
  }

  for _,key := range keys{
    fmt.Printf("Frequency:%v\t Character:%v\n",resulting_map[byte(key)],string(key)) 
    huffman_node_arr = append(huffman_node_arr, huffman_node{resulting_map[byte(key)],key,nil,nil})
  }
  
  return huffman_node_arr;
  // now that I have a huffman node for each value, first I will sort the map with respect to the frequency and create a tree
}
func merge_node(n1* huffman_node,n2* huffman_node) huffman_node{
  return huffman_node{n1.freq+n2.freq,0b00110011,n1,n2}
}
func push_and_sort(huff_node_to_insert huffman_node, huffman_node_arr []huffman_node)[]huffman_node{
  
  var resulting_huff_node_arr []huffman_node
  for index,iter_huff_node := range huffman_node_arr{
    if iter_huff_node.freq>huff_node_to_insert.freq{
      resulting_huff_node_arr = append(huffman_node_arr[0:index],huff_node_to_insert )
      resulting_huff_node_arr = append(resulting_huff_node_arr,huffman_node_arr[index:]...)
      fmt.Printf("[SHORT-MERGE]-%v\n",huffman_node_arr);
      return resulting_huff_node_arr;
    }
  } 
      resulting_huff_node_arr = append(huffman_node_arr,huff_node_to_insert)
      return resulting_huff_node_arr
}

func traverse_root(root_node *huffman_node) *huffman_node{
  if root_node == nil {
    return nil;
  }
  traverse_root(root_node.left_node);
  fmt.Printf("Character:%s\nFrequency:%d\n",string(root_node.char_data),root_node.freq)
  return traverse_root(root_node.right_node);
  return nil;
}
