NAME=exercice2.2

CC=gcc
CFLAGS=-Wall -std=c23

$(NAME): $(NAME).lib.o $(NAME).test.o
	$(CC) $(CFLAGS) $(NAME).lib.o $(NAME).test.o -o $(NAME) 
$(NAME).lib.o: $(NAME).lib.c
	$(CC) $(CFLAGS) -c $(NAME).lib.c
$(NAME).test.o: $(NAME).test.c
	$(CC) $(CFLAGS) -c $(NAME).test.c	 

clean:
	rm -f $(NAME) *.o		
