# Minimal Linux image
FROM alpine:latest

# Install bash (optional) and set terminal
RUN apk add --no-cache bash
ENV TERM xterm-256color

# Set working directory
WORKDIR /app

# Copy your Linux binary and JSON file(s)
COPY mytui .
COPY todos.json .

# Make the binary executable
RUN chmod +x mytui

# Run your TUI
ENTRYPOINT ["./mytui"]
