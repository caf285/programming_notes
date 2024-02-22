public class Primitives {
    public static void primitives() {
        /* wrapper classes have various properties like max and min value */
        System.out.println(Integer.MAX_VALUE);

        /* anything after '+' in a System.out.println is converted to a String */
        System.out.println("the following is converted to a string: " + 3.14);

        /* any amount of literals can be strung together */
        System.out.println("The range of values for int are (" + Integer.MIN_VALUE + " to " + Integer.MAX_VALUE + ")");
    }
}