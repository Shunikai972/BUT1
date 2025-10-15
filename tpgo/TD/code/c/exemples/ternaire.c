//ternaire.c

#include <assert.h>

int abs(int a) {
	//Si a>0 alors a sinon -a
	return a > 0 ? a : -a;
}

//La fonction main 
int main(int argc, char *argv[]){
	assert(abs(2)==2);
	assert(abs(-2)==2);
	assert(abs(0)==0);
	return 0;
}