// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)

// Put your code here.

// SETUP
  @i      // load i into A
  M=1     // set M[i] to 1
  @R2     // load R2 into A
  M=0     // set M[A] to zero
// CHECK LOOP CONDITION
(LOOP)    // create location to jump to
  @i      // load i into A
  D=M     // load M[A] into D
  @R1     // load R1 into A
  D=D-M   // subtract M[A] from D, store in D
  @END    // load M[END]
  D;JGT   // Jump if D greater than zero
// PERFORM CALCULATION
  @R2     // load R2 into A
  D=M     // store M[A] in D
  @R0     // load R0 into A
  D=D+M   // add D and M[A], store in D
  @R2     // load R2 into A
  M=D     // store D into M[R2]
// INCREMENT LOOP
  @i      // load i into A
  M=M+1   // add 1 to M[A], store in M[A]
  @LOOP   // load LOOP into A
  0;JMP   // Jump to Loop
// END
(END)     // create location to jump to
  @END    // create infinite loop
  0;JMP