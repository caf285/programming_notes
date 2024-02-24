public class OperatorsExpressions {
    public static void arithmetic() {
        /* operators are special symbols to perform specific operations and return results */
        /* arithmetic operators include '+', '-', '*', '/', and '%' */
        System.out.println("Addition: 1 + 2 = " + (1 + 2));
        System.out.println("Subtraction: 5 - 20 = " + (5 - 20));
        System.out.println("Multiplication: 2 * 12.5 = " + (2 * 12.5D));
        System.out.println("Division: 5 / 3 = " + (5D / 3D));
        System.out.println("Modulo: 10 % 3 = " + (10 % 3));
    }
    
    public static void unary() {
        /* urnary operators include '+', '-', '++', '--', and '!' */
        System.out.println("+(25) = " + +25);
        System.out.println("-(25) = " + -25);
        System.out.println("!(true) = " + !true);

        int x = 25;
        System.out.println("x = 25; ++x results in " + ++x + ", because ++x is calculated before print.");
        x = 25;
        System.out.println("x = 25; x++ results in " + x++ + ", because x is printed before calculating x++. x is actually " + x + ".");
        x = 25;
        System.out.println("x = 25; --x results in " + --x + ", because --x is calculated before print.");
        x = 25;
        System.out.println("x = 25; x-- results in " + x-- + ", because x is printed before calculating x--. x is actually " + x + ".");
    }

    /* x += y   ==>   x += (type of x) + y*/
}
