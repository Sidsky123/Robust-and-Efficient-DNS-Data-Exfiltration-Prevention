import pandas as pd

HEAVY_DATA_SIZE_MB = 147.6 + 90.3
LIGHT_DATA_SIZE_MB = 20.7 + 62.4

def get_time_of_traffic(df, is_light=True):
    # Input unshuffled data, and calculate the time elapsed between the first and last query
    if is_light:
        light_or_heavy_string = "light"
    else:
        light_or_heavy_string = "heavy"

    df_subset = df[df.origin.isin([light_or_heavy_string + "-benign", light_or_heavy_string + "-attack"])]
    timestamps = pd.to_datetime(df_subset.timestamp)
    return (timestamps.iloc[-1] - timestamps.iloc[0]).seconds


def exfiltration_data_size(t, df, is_light=True):
    # Calculates the lower bound on the amount of data exfiltrated when a model takes total time t to detect malicious queries.
    if is_light:
        light_or_heavy_string = "light"
        per_data_point = LIGHT_DATA_SIZE_MB / df.shape[0]
    else:
        light_or_heavy_string = "heavy"
        per_data_point = HEAVY_DATA_SIZE_MB / df.shape[0]

    df_subset = df[df.origin.isin([light_or_heavy_string + "-benign", light_or_heavy_string + "-attack"])]

    queries_per_second = df_subset.shape[0] / get_time_of_traffic(df_subset, is_light=is_light)
    print(queries_per_second)
    total_exfiltrated_data_size = t * queries_per_second * per_data_point
    return total_exfiltrated_data_size

