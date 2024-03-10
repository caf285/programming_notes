public class Main {
    public static void main(String[] args) {
        String myString = "asidouasdbbggryrggbba;sdlk";
        String bestString = "";
        for (int i = 0; i < myString.length(); i++) {
            // even
            int left = i;
            int right = i + 1;
            String loopString = "";
            while (left >= 0 && right < myString.length() && myString.charAt(left) == myString.charAt(right)) {
                loopString = myString.charAt(left--) + loopString + myString.charAt(right++);
                if (bestString.length() < loopString.length()) {
                    bestString = loopString;
                }
            }
            
            // odd
            loopString = "" + myString.charAt(i);
            left = i - 1;
            right = i + 1;
            while (left >= 0 && right < myString.length() && myString.charAt(left) == myString.charAt(right)) {
                loopString = myString.charAt(left--) + loopString + myString.charAt(right++);
                if (bestString.length() < loopString.length()) {
                    bestString = loopString;
                }
            }
        }
        System.out.println(bestString);
    }
}