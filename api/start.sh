# Startup shell

# Install project dependencies
go get ./...

# Activate the cron on the same container
cd /task

# Set right on cronjob
chmod 755 bobba-cron.txt

# Add log to crontab
touch /var/log/bobba-cron.log

# Add crontab.txt
/usr/bin/crontab bobba-cron.txt

# Run the cron
crond

# Go back to task folder
cd /task

# Start and watch the project
CompileDaemon -command="./data"