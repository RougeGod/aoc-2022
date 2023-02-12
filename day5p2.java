import java.util.Scanner;
import java.io.*;

public class day5p1 {
    public static char[][] stacks = new char[9][55];
    //first number is the stack number, 
    //second number is position in the stack, where position 0 represents the 
    //bottom of the stack. 
        

public static void main(String[] cli) {
    Scanner inStream = null;
    String filename = cli[0];
    String[] splitLine = new String[6];
    int[] instructions = new int[3];

/* Hard coded inout was here. It has been redacted*/
    try {
        inStream = new Scanner(new File(filename));
        
        for (int count = 0; count < 10; count++) {
            inStream.nextLine(); //send the first 10 lines to the shadow realm            
            }    

        while(inStream.hasNextLine()) {
            splitLine = inStream.nextLine().split(" ");
            instructions[0] = Integer.parseInt(splitLine[1]);
            instructions[1] = Integer.parseInt(splitLine[3]);
            instructions[2] = Integer.parseInt(splitLine[5]);
         
                push(stacks[instructions[2] - 1], pop(stacks[instructions[1] - 1], instructions[0]), instructions[0]);
            }
        for (int count = 0; count < 9; count++) {
            System.out.print(stacks[count][findEnd(stacks[count])]);            
            }
        

        }
    catch (FileNotFoundException z) {
        System.out.println("File's not here, loser");        
        }
    finally {
        if (inStream != null) {
            inStream.close();            
            }
        }

    }

public static int findEnd(char[] stack) {
    for (int count = stack.length - 1; count >= 0; count--) {
        if (stack[count] != 0) {
            return count;          
            }            
        }
    return 0;
    }

public static void push(char[] destination, char[] sent, int many) {
    for (int count = 0; count < many; count++) {
        destination[findEnd(destination) + 1] = sent[count];
        //findEnd is recalculated every time so this should be fine
        }
    }

public static char[] pop(char[] poppedFrom, int many) {
    if ((int)poppedFrom[findEnd(poppedFrom)] == 0) {
        System.out.println("uh oh");        
        }
    char[] output = new char[many];
    int startingSpot = (findEnd(poppedFrom) + 1 - many);
    for (int count = 0; count < many; count++) {
    output[count] = poppedFrom[startingSpot + count];
    poppedFrom[startingSpot + count] = '\0';
    } 
    return output; 
    }

}
