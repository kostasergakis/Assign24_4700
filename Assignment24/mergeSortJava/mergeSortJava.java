package Assignment24.mergeSortJava;

import java.util.Arrays;
import java.util.Random;

public class mergeSortJava{
    public static void main(String[] args){
        runTest(50);
    }   


    private static void runTest(int numTests){
        int length = 1000;
        int[] sequence;
        int[] sortedSequence;
        long startTime;
        float totalTime; 
        float[] times= new float[numTests];
        boolean sorted;

        for(int i = 0; i < numTests; i++){
            sequence = generateSequence(length);
            startTime = System.currentTimeMillis();
            sortedSequence = sort(sequence);
            totalTime = ((System.currentTimeMillis() - startTime) / 1000F);
            times[i] = totalTime;
            sorted = testAlgo(sortedSequence);
            if(!sorted){
                System.out.println("Sequence of size: " + length + " failed to sort.");
            }
            length += 2000;
        }
        for(int i = 0; i < times.length; i++){
            System.out.print(times[i] + ", ");
        }
    }
    private static boolean testAlgo(int[] sequence){
        boolean passed = false;
        for(int i = 1; i < sequence.length; i++){
            if(sequence[i] >= sequence[i-1]){
                passed = true;
            } else {
                passed = false;
            }
        }
        return passed;
    }

    private static int[] generateSequence(int size){
        Random rand = new Random();
        int[] sequence = new int[size];
        for (int i = 0; i < sequence.length; i++) {
            sequence[i] = rand.nextInt(1000); 
        }

        return sequence;

    }

    private static int[] sort(int[] sequence){
        if (sequence.length == 1){
            return sequence;
        }
        int mid = (int)(sequence.length / 2);
        int[] leftSeq = Arrays.copyOfRange(sequence, 0, mid);
        int[] rightSeq = Arrays.copyOfRange(sequence, mid, sequence.length);

        // for (int i = 0; i < sequence.length; i++){
        //     if (i < mid) {

        //     }
        // }

        return merge(sort(leftSeq), sort(rightSeq));
    }

    private static int[] merge(int[] l, int[] r){
        int[] result = new int[l.length + r.length];

        int i = 0;
        
        while (l.length > 0 && r.length > 0){
            if (l[0] < r[0]){
                result[i] = l[0];
                l = Arrays.copyOfRange(l, 1, l.length);
            }
            else {
                result[i] = r[0];
                r = Arrays.copyOfRange(r, 1, r.length);
            }
            i++;
        }
        for(int j = 0; j < l.length; j++){
            result[i] = l[j];
            i++;
        }
        for(int j = 0; j < r.length; j++){
            result[i] = r[j];
            i++;
        }

        return result;
    }
}
