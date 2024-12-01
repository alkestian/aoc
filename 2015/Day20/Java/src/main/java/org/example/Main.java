package org.example;

public class Main {
    public static void main(String[] args) {
        int input = 36_000_000;
        System.out.println(part2(input));
    }

    public static int part1(int input) {
        int[] housePresents = new int[1_000_000];
        for (int elf = 1; elf < housePresents.length; elf++) {
            for (int house = elf; house < housePresents.length; house += elf) {
                housePresents[house] += elf * 10;
            }
        }

        for (int house = 1; house < housePresents.length; house++) {
            if (housePresents[house] >= input) {
                return house;
            }
        }

        return -1;
    }

    public static int part2(int input){
        int[] housePresents = new int[1_000_000];
        for (int elf = 1; elf < housePresents.length; elf++) {
            for (int house = elf; house <= (elf * 50) && house < housePresents.length; house += elf) {
                housePresents[house] += elf * 11;
            }
        }

        for (int house = 1; house < housePresents.length; house++) {
            if (housePresents[house] >= input) {
                return house;
            }
        }

        return -1;
    }
}