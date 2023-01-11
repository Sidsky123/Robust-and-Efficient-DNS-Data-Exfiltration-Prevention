The analysis begins with the "Prepare Dataset.ipynb" notebook which uses the CIC Data and prepares two csv files "stateless.csv" and "stateful.csv" over the 5 day period combining benign traffic with heavy and light attacks.

We reimplement the 5 models from the main paper 'Lightweight Hybrid Data Exfiltration using DNS based on Machine Learning by Samaneh Mahdavifar Et al.' and measure the time taken, which is in the files:
- RF Model Reimplementation.ipynb
- LR Model Reimplementation.ipynb
- SVM Model Reimplementation.ipynb
- GNB Model Reimplementation.ipynb
- MLPC Model Reimplementation.ipynb

EDA on the stateless features is in the Stateless_EDA.ipynb file.

Feature engineering was done on the Stateless_EDA.ipynb file and the Text-feature Engineering.ipynb files.

The feature selection is done in Final_LR Model Reimplementation with Feature Selection.ipyb and Final_RF Model Reimplementaion with Feature Selection.ipynb, where the final features were selected based on their performance on the LR and RF models respectively. 

Ensemble models were tested in the ensemble_learning.ipynb file. 

Robustness was tested with other DNS exfiltration tools using the file Other_DNS_Exfiltration_tools.ipynb
Malware queries were generated with framework_pos.go on the credit card numbers from CardBase.csv, to generate the file card_data.csv with the encoded data. These are passed to DNS_Malware_Query_Generation.ipynb to generate the malware queries. The queries were tested on the models in the Malware Generated Queries with Model Results.ipynb

