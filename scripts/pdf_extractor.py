import PyPDF2
import sys

def extract_text_from_pdf(pdf_path):
    with open(pdf_path, 'rb') as file:
        reader = PyPDF2.PdfReader(file)
        text = ''
        for page in reader.pages:
            text += page.extract_text() + '\n'
        return text

if __name__ == '__main__':
    if len(sys.argv) != 2:
        print("Usage: python pdf_extractor.py <pdf_file>")
        sys.exit(1)
    
    pdf_path = sys.argv[1]
    try:
        text = extract_text_from_pdf(pdf_path)
        print(text)
    except Exception as e:
        print(f"Error extracting text from PDF: {e}", file=sys.stderr)
        sys.exit(1)
