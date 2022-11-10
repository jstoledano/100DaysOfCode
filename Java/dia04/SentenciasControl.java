import java.util.Scanner;

class SentenciasControl {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.print("Introduce un número: ");
        int numeroIf = sc.nextInt();
        
        // Usando un if
        if (numeroIf > 0) {
            System.out.println("El número es positivo");
        } else if (numeroIf < 0) {
            System.out.println("El número es negativo");
        } else {
            System.out.println("El número es cero");
        }

        // Usando un bucle while
        int numeroWhile = 0;
        while (numeroWhile < 3) {
            numeroWhile++;
            System.out.println("El valor de la variable es: " + numeroWhile);
        }

        // Usando un bucle do-while
        int numeroDoWhile = 2;
        do {
            numeroDoWhile++;
            System.out.println("El valor de la variable es: " + numeroDoWhile);
        } while (numeroDoWhile < 3);

        // Usando un bucle for
        for (int numeroFor = 0; numeroFor < 3; numeroFor++) {
            System.out.println("El valor de la variable es: " + numeroFor);
        }

        // Usando un switch
        String estacion = "otoño";
        switch (estacion) {
            case "primavera":
                System.out.println("La estación es primavera");
                break;
            case "verano":
                System.out.println("La estación es verano");
                break;
            case "otoño":
                System.out.println("La estación es otoño");
                break;
            case "invierno":
                System.out.println("La estación es invierno");
                break;
            default:
                System.out.println("No es una estación");
        }
    }
}


