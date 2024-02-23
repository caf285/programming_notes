public class HelloWorld {
    public static void printHelloWorld() {
        /* print with `System.out.print` */
        System.out.print("Hello World\n");

        /* print new line \n escape character or with `System.out.println` */
        System.out.print("\n");
        System.out.println("Hello World");
    }

    public static void concatination() {
        /* anything after '+' in a System.out.println is converted to a String */
        System.out.println("the following is converted to a string: " + 3.14);

        /* any amount of literals can be strung together */
        System.out.println("The range of values for int are (" + Integer.MIN_VALUE + " to " + Integer.MAX_VALUE + ")");
    }
}
