#!/usr/bin/env python3
import time

from rpi_ws281x import *

# LED strip configurations
LED_COUNT = 240  # Number of LED pixels
LED_PIN = 18  # GPIO pin connected to the pixels (18 uses PWM!)
LED_FREQ_HZ = 800000  # LED signal frequency in Hz (usually 800KHz)
LED_DMA = 10  # DMA channel to use for generating signal (try 10)
LED_BRIGHTNESS = 255  # Set to 0 for darkest and 255 for brightest
LED_INVERT = False  # True to invert the signal (when using NPN transistor level shift)
LED_CHANNEL = 0  # Set to '1' for GPIOs 13, 19, 41, 45, or 53


class Strip:
    def __init__(self):
        self.strip = Adafruit_NeoPixel(LED_COUNT, LED_PIN, LED_FREQ_HZ, LED_DMA, LED_INVERT, LED_BRIGHTNESS, LED_CHANNEL)
        self.strip.begin()
        print("Strip initalized", flush=True)

    def clear(self):
        for i in range(self.strip.numPixels()):
            self.strip.setPixelColor(i, Color(0, 0, 0))
        self.strip.show()

    def solid_color(self, color, intensity, wait_ms=0):
        """Solid color across display"""
        for i in range(self.strip.numPixels()):
            self.strip.setPixelColor(i, Color(color['r'], color['g'], color['b']))
        self.strip.setBrightness(int(255 * intensity))
        self.strip.show()
        time.sleep(wait_ms / 1000.0)

    def fade(self, color1, color2, intensity1, intensity2=1.0, duration_ms=5000):
        """Fade between two colors over a duration"""
        pixels = self.strip.numPixels()
        t, step, red, green, blue = 100, 0, 0, 0, 0
        while step < 1:
            step += (t / float(duration_ms))
            for i in range(int(min(pixels * step, pixels))):
                red = int(color1['r'] + (color2['r'] - color1['r']) * step)
                green = int(color1['g'] + (color2['g'] - color1['g']) * step)
                blue = int(color1['b'] + (color2['b'] - color1['b']) * step)
                self.strip.setPixelColorRGB(
                    i,
                    red if red > 0 else 0,
                    green if green > 0 else 0,
                    blue if blue > 0 else 0
                )
            brightness = int(255 * (intensity2 - intensity1) * step)
            self.strip.setBrightness(brightness)
            self.strip.show()
            time.sleep(t/1000.0)
        return True

    def wheel(self, pos):
        """Generate rainbow colors across 0-255 positions"""
        if pos < 85:
            return Color(pos * 3, 255 - pos * 3, 0)
        elif pos < 170:
            pos -= 85
            return Color(255 - pos * 3, 0, pos * 3)
        else:
            pos -= 170
            return Color(0, pos * 3, 255 - pos * 3)

    def rainbow(self):
        """Rainbow that fades across all pixels at once"""
        for j in range(256):
            for i in range(self.strip.numPixels()):
                self.strip.setPixelColor(i, self.wheel((i + j) & 255))
            self.strip.show()
            time.sleep(50/1000.0)

    def test(self):
        self.clear()
        print('Fade', flush=True)
        self.fade(
            {'r': 0, 'g': -50, 'b': -120},
            {'r': 255, 'g': 130, 'b': 40},
            0.5,
            1.0,
            3000
        )
        print('Solid color', flush=True)
        self.solid_color({'r': 255, 'g': 0, 'b': 0}, 1, 1000)
        self.solid_color({'r': 0, 'g': 255, 'b': 0}, 1, 1000)
        self.solid_color({'r': 0, 'g': 0, 'b': 255}, 1, 1000)
        self.solid_color({'r': 255, 'g': 255, 'b': 255}, 1, 1000)
        print('Rainbow', flush=True)
        self.rainbow()
        self.clear()
