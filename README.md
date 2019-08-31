# The Elements of Computing Systems

The Elements of Computing Systems: Building a Modern Computer from First Principles by Noam Nisan and Shimon Schocken.

While reading this book I am implementing all the projects using Go. So in the end I should have a nice simulation of a computer built up starting at a NAND gate.

page 100
file:///home/bob/Downloads/The%20Elements%20of%20Computing%20Systems.pdf

## Step 1: Gates

The basic unit is the `nand` function. Using this I have implemented other basic gates: `not`, `and`, `or`, `xor` and `nor`. Using these I have also implementen a `mux` and `dmux`. The multiplexor takes 2 inputs and a selector input, the selector decides whether the gate returns input 1 or input 2. The demultiplexor takes 1 input and a selector input, depending on the selector it returns the input on output channel 1 or 2. These gates are then used to create a `mux4Way`, `mux8way`, `dmux4Way` and `dmux8Way`. All of these gates are implemented using array inputs so that the final computer can be 4-bit, 8-bit and such.

## Step 2: ALU

The next units are the `halfAdder` and `fullAdder`. Together with an `add` and `increment` function we have constructed an `alu`. The ALU has 2 multi bit inputs and 6 flags input. The ALU gives multi bit output and two flags, the first flag is whether the output is 0, and the second flag is whether the output is negative.

## Step 3: Sequential

Next up is creating storage. The basic unit is the `dataFlipFlop`, every tick this struct receives an input and sets the `currentValue` equal to the input. The `dataFlipFlop` is used in the `bit`, the `bit` receives two inputs every tick, depending on the load input the input input is stored in the `dataFlipFlop`. A `register` stores an array of `bits`.

The `RAM8` module has 8 `register`s. Every tick it receives an input and load input for the register at the address input. A `RAM` module also has read function to read the bits stored at the address input. The largest `RAM` module is the `ram16k` struct.

Finally there is a special `count` struct. Every tick it takes an input, increment, load and reset. Usually it will receive an increment input which will increase the stored count by one. Otherwise it can take an input and load input to set the count to that input. If the reset input is given the count is set back to zero.

## Step 4: Computer

Finally we construct the full computer. This consists of a `CPU` which has a data-register, address-register and a counter. Every tick the cpu receives an instruction input which comes from the assembly code, the input from memory which can be the second input for the ALU, and a reset flag which is passed to the counter.

The CPU triggers the ALU, and depending on the instruction will store the output in the D register and/or A register, and finally increment/reset/load the counter.

The other parts of the computer are the memory struct which is a `ram16k` unit, and a `rom32k` struct in which the program is stored. The computer has a `loadProgram` function that takes a filepath and loads that program into the ROM. The file should be `.hack` file with a binary instruction on every line. To run the program the `tick` function of the computer should be called.
