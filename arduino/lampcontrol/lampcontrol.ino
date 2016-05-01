#include <Adafruit_NeoPixel.h>
#ifdef __AVR__
  #include <avr/power.h>
#endif

#define PIN 6
#define PIXELS 4*15 // Total number of LED's connected

// Parameter 1 = number of pixels in strip
// Parameter 2 = Arduino pin number (most are valid)
// Parameter 3 = pixel type flags, add together as needed:
//   NEO_KHZ800  800 KHz bitstream (most NeoPixel products w/WS2812 LEDs)
//   NEO_KHZ400  400 KHz (classic 'v1' (not v2) FLORA pixels, WS2811 drivers)
//   NEO_GRB     Pixels are wired for GRB bitstream (most NeoPixel products)
//   NEO_RGB     Pixels are wired for RGB bitstream (v1 FLORA pixels, not v2)
Adafruit_NeoPixel strip = Adafruit_NeoPixel(PIXELS, PIN, NEO_GRBW + NEO_KHZ800);

// IMPORTANT: To reduce NeoPixel burnout risk, add 1000 uF capacitor across
// pixel power leads, add 300 - 500 Ohm resistor on first pixel's data input
// and minimize distance between Arduino and first pixel.  Avoid connecting
// on a live circuit...if you must, connect GND first.

int PixelCount = PIXELS;
int PixelStartEnd = 0; // Start/End point for Cylon effect

// static colors
uint32_t Red = strip.Color(255, 0, 0, 0);
uint32_t Orange = strip.Color(255, 127, 0, 0);
uint32_t Yellow = strip.Color(255, 255, 0, 0);
uint32_t Green = strip.Color(0, 255, 0, 0);
uint32_t Blue = strip.Color(0, 0, 255, 0);
uint32_t Indigo = strip.Color(75, 0, 130, 0);
uint32_t Violet = strip.Color(138, 43, 226, 0);
uint32_t Purple = strip.Color(127, 0, 127, 0);
uint32_t White = strip.Color(0, 0, 0, 255);

char command[20];

void setup() {
  // This is for Trinket 5V 16MHz, you can remove these three lines if you are not using a Trinket
  #if defined (__AVR_ATtiny85__)
    if (F_CPU == 16000000) clock_prescale_set(clock_div_1);
  #endif
  // End of trinket special code

  delay(2500); // give some time to open serial monitor;
  Serial.begin(250000);
  Serial.println("Starting...");

  strip.begin();
  strip.setBrightness(250);
  strip.show(); // Initialize all pixels to 'off'
}

void loop() {
  while(!Serial.available()); // wait for input
  delay(10);
  
  //Serial.readBytes(&command[0], Serial.available());
  int inByte = Serial.read();
  String color = "";

  switch(inByte) {
    case 'r': // change color to red
      color = "red";
      colorWipe(Red, 0, color);
      break;
    case 'o': // orange
      color = "orange";
      colorWipe(Orange, 0 , color);
      break;
    case 'y': // yellow
      color = "yellow";
      colorWipe(Yellow, 0, color);
      break;
    case 'g': // change color to green
      color = "green";
      colorWipe(Green, 0, color);
      break;
    case 'b': // change color to blue
      color = "blue";
      colorWipe(Blue, 0, color);
      break;
    case 'i': // indigo
      color = "indigo";
      colorWipe(Indigo, 0, color);
      break;
    case 'v': // Violet
      color = "violet";
      colorWipe(Violet, 0, color);
      break;
    case 'p': // purple   
      color = "purple";
      colorWipe(Purple, 0, color);
      break;
    case 'w': // change color to white
      color = "white";
      colorWipe(White, 0, color);
      break;
    case 'q': // turn the strip off
      clearStrip();
      break;
    case 'c': // act like a cylon
      Cylon(); // TODO:: Teach this thing to handle interrupts...
      break;
    case '1': // ColorTest
      colorWipe(Red, 20, "red"); // Red
      colorWipe(Green, 20, "green"); // Green
      colorWipe(Blue, 20, "blue"); // Blue
      colorWipe(White, 20, "white"); // White
      clearStrip();
      delay(2000);
      colorWipe(White, 20, "white"); // White
      break;
    case '2': // TheaterChase
      // Send a theater pixel chase in...
      theaterChase(Red, 100); // Red
      theaterChase(Green, 100); // Green
      theaterChase(Blue, 100); // Blue
      theaterChase(White, 100); // White
      clearStrip();
      delay(2000);
      colorWipe(White, 20, "white"); // White
      break;
    case '3': // Rainbow
      rainbow(20);
      break;
    case '4': // Rainbow Cycle
      rainbowCycle(20);
      break;
    case '5': // TheaterChaseRainbow
      theaterChaseRainbow(100);
      break;
    default:
      Serial.println("Got " + inByte);
      clearStrip();
      break;
  }
}

void resetCommandBuffer() {
  for (int i=0; i<20; i++) {
    command[i] = 0;
  }
}

void clearStrip() {
  Serial.println("Turning LEDs off.");
  Serial.flush();
  for(uint16_t i=0; i<strip.numPixels(); i++) {
    strip.setPixelColor(i, strip.Color(0,0,0,0));
    strip.show();
  }
}

// Fill the dots one after the other with a color
void colorWipe(uint32_t c, uint8_t wait, String color) {
  Serial.println("Colorwipe " + color);
  Serial.flush();
  for(uint16_t i=0; i<strip.numPixels(); i++) {
    strip.setPixelColor(i, c);
    strip.show();
    delay(wait);
  }
}

void rainbow(uint8_t wait) {
  uint16_t i, j;

  for(j=0; j<256; j++) {
    for(i=0; i<strip.numPixels(); i++) {
      strip.setPixelColor(i, Wheel((i+j) & 255));
    }
    strip.show();
    delay(wait);
  }
}

// Slightly different, this makes the rainbow equally distributed throughout
void rainbowCycle(uint8_t wait) {
  uint16_t i, j;

  for(j=0; j<256*5; j++) { // 5 cycles of all colors on wheel
    for(i=0; i< strip.numPixels(); i++) {
      strip.setPixelColor(i, Wheel(((i * 256 / strip.numPixels()) + j) & 255));
    }
    strip.show();
    delay(wait);
  }
}

//Theatre-style crawling lights.
void theaterChase(uint32_t c, uint8_t wait) {
  for (int j=0; j<10; j++) {  //do 10 cycles of chasing
    for (int q=0; q < 3; q++) {
      for (int i=0; i < strip.numPixels(); i=i+3) {
        strip.setPixelColor(i+q, c);    //turn every third pixel on
      }
      strip.show();

      delay(wait);

      for (int i=0; i < strip.numPixels(); i=i+3) {
        strip.setPixelColor(i+q, 0);        //turn every third pixel off
      }
    }
  }
}

//Theatre-style crawling lights with rainbow effect
void theaterChaseRainbow(uint8_t wait) {
  for (int j=0; j < 256; j++) {     // cycle all 256 colors in the wheel
    for (int q=0; q < 3; q++) {
      for (int i=0; i < strip.numPixels(); i=i+3) {
        strip.setPixelColor(i+q, Wheel( (i+j) % 255));    //turn every third pixel on
      }
      strip.show();

      delay(wait);

      for (int i=0; i < strip.numPixels(); i=i+3) {
        strip.setPixelColor(i+q, 0);        //turn every third pixel off
      }
    }
  }
}

// Input a value 0 to 255 to get a color value.
// The colours are a transition r - g - b - back to r.
uint32_t Wheel(byte WheelPos) {
  WheelPos = 255 - WheelPos;
  if(WheelPos < 85) {
    return strip.Color(255 - WheelPos * 3, 0, WheelPos * 3);
  }
  if(WheelPos < 170) {
    WheelPos -= 85;
    return strip.Color(0, WheelPos * 3, 255 - WheelPos * 3);
  }
  WheelPos -= 170;
  return strip.Color(WheelPos * 3, 255 - WheelPos * 3, 0);
}

void Cylon() {
  // clear the strip from whatever was happening before...
  clearStrip();
  Serial.println("Starting Cylon...");
  Serial.flush();
  int wait_T = 40;
  uint32_t color1 = strip.Color(255,0,0,0);
  uint32_t color2 = strip.Color(100,0,0,0);
  uint32_t color3 = strip.Color(10,0,0,0);
  
  while(true) {
    // Break if serial input is waiting...
    if (Serial.available() > 0) {
      int brk = Serial.read();
      // ignore line feeds...
      if (brk != 10) {
        Serial.print(brk);
        Serial.flush();
        break;  
      }
    }
    //Example: CylonEyeUp(Center_Dot_Color, Second_Dot_color, Third_Dot_color, wait_T, PixelCount, Pixel_Start_End);
    CylonEyeUp(color1, color2, color3, wait_T, PixelCount, PixelStartEnd);
    delay(wait_T);
    //Example: CylonEyeDown(Center_Dot_Color, Second_Dot_color, Third_Dot_color, wait_T, PixelCount, Pixel_Start_End);
    CylonEyeDown(color1, color2, color3, wait_T, PixelCount, PixelStartEnd);
    delay(wait_T);
  }
  Serial.println("Exiting Cylon.");
  Serial.flush();
}

void CylonEyeUp(uint32_t Co, uint32_t Ct, uint32_t Ctt, uint8_t Delay, int TotalPixels, int pStart) {
  for(int i=pStart; i<TotalPixels; i++) {
    //if(!UsingBar) { 
    strip.setPixelColor(i+2, Ctt); //Third Dot Color
    //} 
    strip.setPixelColor(i+1, Ct);   //Second Dot Color
    strip.setPixelColor(i, Co);     //Center Dot Color
    strip.setPixelColor(i-1, Ct);   //Second Dot Color
    //if(!UsingBar) { 
    strip.setPixelColor(i-2, Ctt); //Third Dot Color
    //} 

    //if(!UsingBar) {
    strip.setPixelColor(i-3, strip.Color(0,0,0)); //Clears the dots after the 3rd color
    //} else {
    //  strip.setPixelColor(i-2, strip.Color(0,0,0)); //Clears the dots after the 2rd color
    //}
    strip.show();
    //Serial.println(i); //Used For pixel Count Debugging
    delay(Delay);
  }
}
void CylonEyeDown(uint32_t Co, uint32_t Ct, uint32_t Ctt, uint8_t Delay, int TotalPixels, int pEnd) {
  for(int i=TotalPixels-1; i>pEnd; i--) {
    //if(!UsingBar) { 
    strip.setPixelColor(i-2, Ctt); //Third Dot Color
    //} 
    strip.setPixelColor(i-1, Ct);   //Second Dot Color
    strip.setPixelColor(i, Co);    //Center Dot Color
    strip.setPixelColor(i+1, Ct);  //Second Dot Color
    // if(!UsingBar) { 
    strip.setPixelColor(i+2, Ctt); //Third Dot Color
    //} 

    //if(!UsingBar) { 
    strip.setPixelColor(i+3, strip.Color(0,0,0)); //Clears the dots after the 3rd color
    //} else {
    //  strip.setPixelColor(i+2, strip.Color(0,0,0)); //Clears the dots after the 2rd color
    //}
    strip.show();
    //Serial.println(i); //Used For pixel Count Debugging
    delay(Delay);
  }
}

