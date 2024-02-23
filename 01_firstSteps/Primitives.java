public class Primitives {
    public static void integerWrapper() {
        /* wrapper classes have various properties like max and min value */
        System.out.println(Integer.MAX_VALUE);

        /* anything after '+' in a System.out.println is converted to a String */
        System.out.println("the following is converted to a string: " + 3.14);

        /* any amount of literals can be strung together */
        /* int occupies 32 bits with values between -2147483648 and 2147483647 */
        System.out.println("The range of values for int are (" + Integer.MIN_VALUE + " to " + Integer.MAX_VALUE + ")");
    }

    public static void byteWrapper() {
        /* byte occupies 8 bits and ranges between -128 and 127 */
        System.out.println("The range of values for int are (" + Byte.MIN_VALUE + " to " + Byte.MAX_VALUE + ")");
    }

    public static void shortWrapper() {
        /* short occupies 16 bits and ranges between -32768 and 32767 */
        System.out.println("The range of values for int are (" + Short.MIN_VALUE + " to " + Short.MAX_VALUE + ")");
    }

    public static void longWrapper() {
        /* long occupies 64 bits and is the largest between -9223372036854775808 to 9223372036854775807 */
        System.out.println("The range of values for int are (" + Long.MIN_VALUE + " to " + Long.MAX_VALUE + ")");

        /* add an 'L' or 'l' suffix to the end of a number literal to denote a long */
        /* any literal larger than int max size must have L to denote long */
        long myLong = 100L;
        System.out.println("myLong value is " + myLong);
    }

    public static void challenge() {
        /* create a byte, short, and int of any values. Next create a long that is equal to 50,000 + 10 times the sum of the first 3 values */
        byte myByte = 10;
        short myShort = 20;
        int myInt = 50;
        long longResult = 50000L + 10L * (myByte + myShort + myInt);
        System.out.println("The challenge answer is " + longResult);

        /* casting required if assigning calculation into a short */
        short shortResult = (short) (10 * (myByte + myShort + myInt));
        System.out.println("Casting short type challenge result is " + shortResult);
    }

    public static void floatWrapper() {
        /* float occupies 32 bits */
        System.out.println("The range of values for a float are (" + Float.MAX_VALUE + " to " + Float.MIN_VALUE + ")");

        /* add an 'F' or 'f' suffix to the end of a number literal to denote a float */
        /* casting or F suffix is required when assigning a rational number literal to a type float */
        float myFloat = 5.25F; // <--more preferred method
        myFloat = (float) 5.25;
        System.out.println("myFloat value is " + myFloat);
    }

    public static void doubleWrapper() {
        /* double occupies 64 bits, is much more precise than a float, and is also faster to process than a double */
        System.out.println("The range of values for a double are (" + Double.MAX_VALUE + " to " + Double.MIN_VALUE + ")");
        
        /* add a 'D' or 'd' suffix to the end of a number literal to denote a double */
        /* D suffix is not required, because Double is the default rational type */
        double myDouble = 5.25D;

        /* when performing calculations, the literals must contain decimals or containd a D suffix on at least one value */
        myDouble = 5D / 3D;
        myDouble = 5.0 / 3.0; // <--more preferred method
        myDouble = 5.0D / 3.0D;
        System.out.println("myDouble value is " + myDouble);
    }

    public static void charWrapper() {
        /* char occupies 16 bits representing unicode characters */
        /* char requires 'single quotes' for characters, but also accepts unicode */
        char myChar = 'A'; // char
        myChar = '\u0041'; // unicode
        myChar = 65; // html code
        System.out.println(myChar);
        System.out.println("The range of values for a char are (" + Character.MAX_VALUE + " to " + Character.MIN_VALUE + ")");
    }

    public static void booleanWrapper() {
        /* booleans are either true or false */
        boolean myBoolean = true;
        System.out.println(myBoolean);

        /* naming convention should be a question */
        boolean isProgrammer = true;
        System.out.print("Am I a programmer? " + isProgrammer);
        System.out.println("The range of values for a boolean are (" + Boolean.TRUE + " to " + Boolean.FALSE + ")");
    }
}