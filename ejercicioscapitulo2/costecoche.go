// fichero costecoche.cpp 
// Este programa calcula el precio total de un coche

#include <iostream>
#include <cstdlib>
#include <iomanip>
using namespace std;
const double ganancia = 1.15;
const double impuesto = 1.21;

int main() {

   double costecoche,pvantesimpuestos,preciototal;    

   cout<<"Introduce el coste del vehiculo: ";
   if (!(cin >> costecoche)) {
        cout << "Error: El valor debe de ser un número.\n";
        return EXIT_FAILURE; 
   } 
   if (costecoche < 0.0) {
        cout << "Error: El valor debe ser un número positivo\n";
        return EXIT_FAILURE;
   }
   pvantesimpuestos= costecoche*ganancia;
   preciototal= pvantesimpuestos*impuesto;
   cout << fixed << setprecision(2);
   cout<<"El precio total del automovil nuevo para el comprador es:"<< preciototal <<"\n";
   return EXIT_SUCCESS;
}
