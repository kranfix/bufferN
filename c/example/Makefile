exe=example.run

${exe}: main.c imgdf.o
	gcc $^ -o ${exe}

run: ${exe}
	./${exe}

imgdf.o: ../imgdf.c ../imgdf.h
	gcc -c $<

clean:
	rm example.run
