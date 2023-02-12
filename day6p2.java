import java.util.Scanner;
import java.io.*;
import java.util.Arrays;

public class Elf {


public static void main(String[] cli) {
    Scanner inStream = null;

    try {
        inStream = new Scanner(new File(cli[0]));    
        String input = inStream.next();
        char[] last14 = new char[14];

        int count = 0;
        while (count < input.length() - 13) {
            for (int loop = 0; loop < 14; loop++) {
                last14[loop] = input.charAt(count + loop);                   
            } 
            if (checkW(last14)) {
                System.out.println(count + 14);
                System.exit(0);                
                }            
            count++;
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
    //sort the last 14 chars so easier to check for duplicates
    public static boolean checkW(char[] freqcheq) {
        Arrays.sort(freqcheq);
        for (int count = 0; count < freqcheq.length - 1; count++) {
            if (freqcheq[count] == freqcheq[count + 1]) {
                System.out.println(Arrays.toString(freqcheq));
                return false;                
                }
            }
        return true;
        }

}
