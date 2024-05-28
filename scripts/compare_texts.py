import sys
import os
import logging
import json
from flair.embeddings import TransformerDocumentEmbeddings
from flair.data import Sentence
from scipy.spatial.distance import cosine

logging.basicConfig(level=logging.INFO)

def load_and_embed(text):
    # Load the pre-trained model
    model = TransformerDocumentEmbeddings('bert-base-uncased')
    # Create a Flair Sentence object
    sentence = Sentence(text)
    # Embed the text using the model
    model.embed(sentence)
    return sentence.embedding

def calculate_similarity(embedding1, embedding2):
    # Calculate cosine similarity (convert to cosine distance to similarity)
    return 1 - cosine(embedding1, embedding2)

def compare_texts(claim, article):
    #logging.info("Start comparing")

    # Load and embed both articles
    #logging.info("Loading and embedding claim")
    embedding1 = load_and_embed(claim)
    #logging.info("Loading and embedding article")
    embedding2 = load_and_embed(article)
    
    # Calculate and print the similarity
    #logging.info("Calculating similarity")
    similarity = calculate_similarity(embedding1, embedding2)
    #logging.info(f"Similarity calculated: {similarity}")
    return(json.dumps({"similarity": similarity}))

if __name__ == "__main__":
    sys.stdout = open(os.devnull, 'w')
    sys.stderr = open(os.devnull, 'w')
    claim = sys.argv[1]
    article = sys.argv[2]
    similarity_score = compare_texts(claim, article)
    sys.stdout = sys.__stdout__
    print(similarity_score)

