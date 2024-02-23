public class Variables {
    public static void declaration() {
        /* variables are declared with a type and variable name */
        int myInteger = 5;
        System.out.println(myInteger);

        /* variable values can be changed without addressing a type after declaration */
        myInteger = 10;
        System.out.println(myInteger);

        /* variables can be used to assign new values */
        int nextInteger = myInteger * 2;
        System.out.println(nextInteger);

        /* multiple variables of the same data type may be declared in java */
        int firstValue = 1, secondValue = 2;
        System.out.println(firstValue + secondValue);
    }

    public static void casting() {
        /* when using larger types for calculations before returning to smaller types, casting is required */
        byte byte1 = 10;
        byte byte2 = 20;
        int resultInteger = byte1 + byte2; // result is promoted to an int
        byte resultByte = (byte) (byte1 + byte2); // casting is required to demote back to byte
        System.out.println("resulting casting calculation: " + resultInteger + ", " + resultByte);
    }
}
