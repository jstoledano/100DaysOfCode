/* Program to calculate the product of two numbers  */

#include <stdio.h>

int val1, val2, val3;

int product( int x, int y );

int main( void ) {
  /* Get the first number */
  printf( "Enter a number between 1 and 100: " );
  scanf( "%d", &val1 );

  /* Get the second number */
  printf( "Enter another number between 1 an 100: " );
  scanf( "%d", &val2 );

  /* Calculate and display the product */
  val3 = product( val1, val2 );
  printf( "%d times %d = %d\n", val1, val2, val3);

  return 0;
}

/* Function return the product of the two values provided */
int product ( int x, int y ) {
  return ( x * y );
}


