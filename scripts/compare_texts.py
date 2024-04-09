from flair.embeddings import TransformerWordEmbeddings
from flair.data import Sentence
import sys

def compare_texts(claim, article):
    # Load pre-trained transformer word embeddings
    embeddings = TransformerWordEmbeddings('bert-base-uncased')

    # Create Sentence objects for the claim and article
    claim_sentence = Sentence(claim)
    article_sentence = Sentence(article)

    # Embed the sentences
    embeddings.embed(claim_sentence)
    embeddings.embed(article_sentence)

    # Compute cosine similarity between the sentence embeddings
    similarity = claim_sentence.get_embedding().similarity(article_sentence.get_embedding())

    return similarity

if __name__ == "__main__":
    claim = sys.argv[1]
    article = sys.argv[2]
    similarity_score = compare_texts(claim, article)
    print(similarity_score)
