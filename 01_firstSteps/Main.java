public class Main {
    public static void main(String[] args) {
        /* 01. hello world */
        HelloWorld.printHelloWorld();
        HelloWorld.concatination();

        /* 02. variables */
        Variables.declaration();
        Variables.casting();

        /* 03. primitives */
        Primitives.integerWrapper();
        Primitives.byteWrapper();
        Primitives.shortWrapper();
        Primitives.longWrapper();
        Primitives.challenge();
        Primitives.floatWrapper();
        Primitives.doubleWrapper();
        Primitives.charWrapper();
        Primitives.booleanWrapper();

        /* 04. built in classes */
        BuiltInClasses.string();

        /* 05. operators, operands, and expressions */
        OperatorsExpressions.arithmetic();
        OperatorsExpressions.unary();
    }
}
