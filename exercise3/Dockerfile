FROM python
WORKDIR /starwarsapifetch
RUN pip install PyYAML
COPY . .
ENTRYPOINT ["python", "starwarscronjob.py"]
CMD ["-i", "input.yaml", "-o", "."]