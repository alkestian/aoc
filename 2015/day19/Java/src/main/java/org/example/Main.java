package org.example;

import java.io.BufferedReader;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.util.ArrayList;
import java.util.Comparator;
import java.util.HashSet;
import java.util.List;
import java.util.stream.Stream;

public class Main {
    public static void main(String[] args) {
        String moleculeString = "CRnSiRnCaPTiMgYCaPTiRnFArSiThFArCaSiThSiThPBCaCaSiRnSiRnTiTiMgArPBCaPMgYPTiRnFArFArCaSiRnBPMgArPRnCaPTiRnFArCaSiThCaCaFArPBCaCaPTiTiRnFArCaSiRnSiAlYSiThRnFArArCaSiRnBFArCaCaSiRnSiThCaCaCaFYCaPTiBCaSiThCaSiThPMgArSiRnCaPBFYCaCaFArCaCaCaCaSiThCaSiRnPRnFArPBSiThPRnFArSiRnMgArCaFYFArCaSiRnSiAlArTiTiTiTiTiTiTiRnPMgArPTiTiTiBSiRnSiAlArTiTiRnPMgArCaFYBPBPTiRnSiRnMgArSiThCaFArCaSiThFArPRnFArCaSiRnTiBSiThSiRnSiAlYCaFArPRnFArSiThCaFArCaCaSiThCaCaCaSiRnPRnCaFArFYPMgArCaPBCaPBSiRnFYPBCaFArCaSiAl";

        try (InputStream inputStream = Main.class.getResourceAsStream("/input.txt");
             BufferedReader reader = new BufferedReader(new InputStreamReader(inputStream))) {

            Stream<String> content = reader.lines();

            //HashSet<String> distinctCombos = part1(content, moleculeString);

            int stepsToE = part2(content, moleculeString);

            System.out.println(stepsToE);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    public static HashSet<String> part1(Stream<String> content, String moleculeString) {
        HashSet<String> distinctCombos = new HashSet<>();

        content.forEach(line -> {
            String oldPiece = line.split("=>")[0].trim();
            String newPiece = line.split("=>")[1].trim();

            int startIndex = moleculeString.indexOf(oldPiece);
            while (startIndex != -1) {
                String newMolecule = moleculeString.substring(0, startIndex) + newPiece + moleculeString.substring(startIndex + oldPiece.length());
                distinctCombos.add(newMolecule);

                startIndex = moleculeString.indexOf(oldPiece, startIndex + 1);
            }
        });

        return distinctCombos;
    };

    public static int part2(Stream<String> content, String moleculeString) {
        List<String[]> replacements = new ArrayList<>();
        content.forEach(line -> {
            replacements.add(new String[]{line.split("=>")[1].trim(), line.split("=>")[0].trim()});
        });
        replacements.sort(Comparator.comparingInt((String[] r) -> r[0].length()).reversed());

        int steps = 0;
        while (!moleculeString.equals("e")) {
            for (String[] replacement : replacements) {
                String target = replacement[0];
                String replacementSource = replacement[1];

                int index = moleculeString.indexOf(target);
                if (index != -1) {
                    moleculeString = moleculeString.substring(0, index) + replacementSource + moleculeString.substring(index + target.length());
                    steps++;
                    break;
                }
            }
        }

        return steps;
     }
}