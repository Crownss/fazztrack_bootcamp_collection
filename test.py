import os

def main(path)->os.system:
    test_cov = "cd ~/programming/fazztrack/hardskill/week_10/go/src/modules/v1/users && go test -coverprofile=coverage.out&& go tool cover -html=coverage.out"
    
    command_herokulogs="heroku logs --tail"
    command_k6="k6 run "+path
    return os.system(test_cov), os.system("gnome-terminal -e 'bash -c \""+command_herokulogs+";bash\"'"), os.system("gnome-terminal -e 'bash -c \""+command_k6+";bash\"'")

if __name__ == '__main__':
    main("k6_test/index.js")