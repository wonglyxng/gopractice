package main

/*
#include <stdio.h>
#include <stdlib.h>

typedef struct {
	float x,y;
	char *name;
	int height;
	int weight;
} Role;

int printRole(){
	int a,b,c;
	a = 4;
	b = 5;
	c = a+b;
	return c;
}

 */
import "C"
import "fmt"

func main() {
	fmt.Println("size=",C.sizeof_Role)
}
