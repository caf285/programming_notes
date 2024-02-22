public class Variables {
    public static void printVariables() {
        /* variables are declared with a type and variable name */
        int myInteger = 5;
        System.out.println(myInteger);

        /* variable values can be changed without addressing a type after declaration */
        myInteger = 10;
        System.out.println(myInteger);

        /* variables can be used to assign new values */
        var nextInteger = myInteger * 2;
        System.out.println(nextInteger);
    }
}
