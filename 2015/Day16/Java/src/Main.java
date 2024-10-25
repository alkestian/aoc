import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.util.HashMap;
import java.util.Map;

@SuppressWarnings("ALL")
public class Main {
    public static void main(String[] args) throws IOException {
        HashMap<String, HashMap<String, Integer>> theSues = new HashMap<String, HashMap<String, Integer>>();
        InputStream input = Main.class.getResourceAsStream("/input.txt");
        BufferedReader reader = new BufferedReader(new InputStreamReader(input));
        String line;
        while ((line = reader.readLine()) != null) {
            String[] pieces = line.split(":");
            String sue = pieces[0].trim();
            StringBuilder attributesBuilder = new StringBuilder();
            for (int i = 1; i < pieces.length; i++) {
                attributesBuilder.append(pieces[i]);
                if (i < pieces.length - 1) {
                    attributesBuilder.append(":");
                }
            }

            String[] attributes = attributesBuilder.toString().split(",");

            HashMap<String, Integer> thisSueAtt = new HashMap<>();

            for (String attribute : attributes) {
                String[] parts = attribute.split(":");

                if (parts.length == 2) {
                    String key = parts[0].trim();
                    int value = Integer.parseInt(parts[1].trim());
                    thisSueAtt.put(key, value);
                } else {
                    System.out.println("Invalid attribute format: " + attribute);
                }
            }

            theSues.put(sue, thisSueAtt);
        }


        HashMap<String, Integer> suePrime = new HashMap<String, Integer>() {{
                put("children", 3);
                put("cats", 7);
                put("samoyeds", 2);
                put("pomeranians", 3);
                put("akitas", 0);
                put("vizslas", 0);
                put("goldfish", 5);
                put("trees", 3);
                put("cars", 2);
                put("perfumes", 1);
        }};

        HashMap<String, Integer> sueCount = new HashMap<>();

        for (String sue : theSues.keySet()) {
            HashMap<String, Integer> thisSueAtt = theSues.get(sue);
            int matchCounter = 0;
            for (String attribute : thisSueAtt.keySet()) {
                if (attribute.equals("cats") || attribute.equals("trees")) {
                    if (thisSueAtt.get(attribute) > suePrime.get(attribute)) {
                        matchCounter++;
                    }
                } else if (attribute.equals("pomeranians") || attribute.equals("goldfish")) {
                    if (thisSueAtt.get(attribute) < suePrime.get(attribute)) {
                        matchCounter++;
                    }
                } else {
                    if (thisSueAtt.get(attribute) == suePrime.get(attribute)) {
                        matchCounter++;
                    }
                }
            }
            sueCount.put(sue, matchCounter);
        }

        int maxValue = 0;
        String theTrueSue = null;

        for (Map.Entry<String, Integer> entry : sueCount.entrySet()) {
            if (entry.getValue() > maxValue) {
                maxValue = entry.getValue();
                theTrueSue = entry.getKey();
            }
        }

        System.out.println(theTrueSue);
    }
}