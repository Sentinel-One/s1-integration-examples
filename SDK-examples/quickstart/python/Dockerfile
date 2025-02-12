# Use an official Python base image
FROM python:3.9-slim

# Set working directory inside the container
WORKDIR /app

# Copy the SentinelDFI wheel file to the container
#COPY nexus-sdk/SDK/SentinelDFI-2-py2.py3-none-any.whl /app/
COPY quickstart/python/scanFile.py /app/
COPY quickstart/python/run.sh /app/
COPY quickstart/sample_files/NexusSDK.pdf /app/
COPY quickstart/sample_files/*.txt /app/

# Make the run.sh script executable
RUN chmod +x /app/run.sh

# Install the wheel file using pip
#RUN python -m pip install SentinelDFI-2-py2.py3-none-any.whl

ENV FILES false
ENV RECURSE false

# Specify the default command to run when the container starts
CMD ["./run.sh"]
