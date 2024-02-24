public class SwitchStatement {
    public static void basicIfElse() {
        /* basic if else structure */
        int value = 2;
        if (value == 1) {
            System.out.println("The value is 1");
        } else if (value == 2) {
            System.out.println("The value is 2");
        }
    }

    public static void basicSwitch() {
        /* switch allows for many checks with optional default basecase */
        /* switch ONLY handles byte, short, int, char, Byte, Short, Integer, Character, String, and enum */
        int value = 10;
        switch(value) {
            case 1:
                System.out.println("The value is 1");
                break;
            case 2:
                System.out.println("The value is 2");
                break;
            case 3: case 4: case 5:
                System.out.println("The value was a 3, 4, or 5");
                System.out.println("The actual value was " + value);
                break;
            default:
                System.out.println("The value is not 1, 2, 3, 4, or 5");
                break;
        }

        /* ommiting break will "fall through" and continue executing cases without checking the case */
        switch (value) {
            case 10:
                System.out.println("The value actually is 10");
            case 11:
                System.out.println("The value is actually 10, but I will execute after fallthrough");
            default:
                System.out.println("... and so will I");
        }
    }

    public static void enhancedSwitch() {
        /* modern java allows for new switch syntax */
        int value = 2;
        switch (value) {
            case 1 -> System.out.println("The value is 1");
            case 2, 3, 4 -> System.out.println("The value is 2, 3, or 4");
            default -> {
                System.out.println("curly braces allow multiple lines");
                System.out.println("I am an additional line");
            }
        }
    }

    public static String getQuarter(String month) {
        /* return statements may be embedded in switch */
        switch (month) {
            case "January":
            case "Februrary":
            case "March":
                return "1st";
            case "April":
            case "May":
            case "June":
                return "2nd";
            case "July":
            case "August":
            case "September":
                return "3rd";
            case "October":
            case "November":
            case "December":
                return "4th";
        }
        return "bad";
    }

    public static String enhancedGetQuarter(String month) {
        /* larger default block requires "yield" keyword rather than "return" */
        return switch (month) {
            case "January", "Februrary", "March" -> "1st";
            case "April", "May", "June" -> "2nd";
            case "July", "August", "September" -> "3rd";
            case "October", "November", "December" -> "4th";
            default -> {
                String badResponse = month + " is bad";
                yield badResponse;
            }
        };
    }
}
