import os
import sys
import json
import statistics
import requests

from google.cloud import vision

os.environ["GOOGLE_APPLICATION_CREDENTIALS"] = "/home/chad/.config/gcloud/application_default_credentials.json"
os.environ["GCLOUD_PROJECT"] = "junction"

class Text_Detector:
    def __init__(self, path, ipaddress, port):
        """Initialisation."""
        self.path = path
        self.tolerance = 7
        self.socket = (ipaddress + ':' + port)
        self.content = {"manufacturer": {},
                        "year": {},
                        "model": {},
                        "serialnum": {}
                        }

    def detect_text(self):
        """Detects text in the file."""

        client = vision.ImageAnnotatorClient()

        with open(self.path, "rb") as image_file:
            content = image_file.read()

        image = vision.Image(content=content)

        response = client.text_detection(image=image)
        if response.error.message:
            raise Exception(
                "{}\nFor more info on error messages, check: "
                "https://cloud.google.com/apis/design/errors".format(response.error.message)
            )
            return
        texts = response.text_annotations
        content = {}
        tmp = 0
        for text in texts:
            if tmp == 0:
                tmp = 1
                continue
            avg = statistics.mean([
                vertex.y for vertex in text.bounding_poly.vertices
            ])
            content[text.description] = avg
        print(content)
        sorted_content = dict(
            sorted(content.items(), key=lambda item: item[1], reverse=False)
        )

        print(sorted_content)
        row_dict = {}
        current_row = []
        current_row_start = None

        for text, position in sorted_content.items():
            if current_row_start is None:
                current_row_start = position
                current_row.append(text)
            elif abs(position - current_row_start) <= self.tolerance:
                current_row.append(text)
            else:
                row_index = len(row_dict) + 1
                row_dict[row_index] = current_row
                current_row = [text]
                current_row_start = position

        if current_row:
            row_index = len(row_dict) + 1
            row_dict[row_index] = current_row

        for key, value in row_dict.items():
            print("row:", key, "value:", value)
            for item in value:
                if "vuosi" in item:
                    for item in value:
                        if item.isnumeric():
                            self.content["year"] = int(item)        
                if "Malli" in item or "Tyyppi" in item:
                    self.content["model"] = ','.join(map(str, value)) 
                if "numero" in item:
                    for item in value:
                        if item.isnumeric():
                            self.content["serialnum"] = item
                if "Valmistaja" in item:
                    self.content["manufacturer"] = ','.join(map(str, value))

    def send_to_server(self, content) -> None:
        """Send JSON to server."""
        print("Sending to server")
        url = 'http://' + self.socket + "/posts"
        myobj = {
            "manufacturer": self.content['manufacturer'],
            "model": self.content['model'],
            "year": self.content['year'],
            "serialnum": self.content['serialnum']
        }
        x = requests.post(url, json = myobj)
        print("Sent:", x.text)
        

def main() -> None:
    text_detector = Text_Detector(sys.argv[1], '192.168.210.213', '8080')
    text_detector.detect_text()
    text_detector.send_to_server(text_detector.content)

if __name__ == '__main__':
    main()
