import java.util.Scanner;
import java.io.*;
import java.util.Arrays;

public class Elf {


public static void main(String[] cli) {
    Scanner inStream = null;

    try {
        inStream = new Scanner(new File(cli[0]));    
        String input = inStream.next();
        
        int count = 0;
        while (count < input.length() - 13) {
            if ((checkEquality(input.charAt(count), input.charAt(count + 1), input.charAt(count + 2), input.charAt(count + 3)) == -1)) {
                //we win if it's -1
                System.out.println(count);
                System.out.print(input.charAt(count)); 
                System.out.print(input.charAt(count + 1));
                System.out.print(input.charAt(count + 2));
                System.out.print(input.charAt(count + 3));

                System.exit(0);                
                }    
            count += checkEquality(input.charAt(count), input.charAt(count + 1), input.charAt(count + 2), input.charAt(count + 3));
            }

        
        }    
    catch (FileNotFoundException x) {
        System.out.println("type the file better, loser");        
        }
    finally {
        if (inStream == null) {
            inStream.close();            
            }        
        }
    
    }

    //returns the number of spaces that the check can advance
    public static int checkEquality(char one, char two, char three, char four) {
        if (three == four) {
            return 3;            
            }
        if ((two == four)||(two == three)) {
            return 2;            
            }
        if (((one == two)||(one == three))||(one == four)) {
            return 1;            
            }
        return -1; //the win condition
        }

}
