


# Teste direto no container

docker run --rm -e CODE="console.log('ok')" ts-runner

Saída esperada:
ok


# Teste via API

curl -X POST http://localhost:8080/run \
 -H "Content-Type: application/json" \
 -d '{"code": "console.log(1 + 2)"}'

Resposta esperada:

{
"stdout": "3\n",
"stderr": "",
"error": ""
}


# Teste de erro

curl -X POST http://localhost:8080/run \
 -H "Content-Type: application/json" \
 -d '{"code": "throw new Error(\"fail\")"}'


# Teste de timeout

curl -X POST http://localhost:8080/run \
 -H "Content-Type: application/json" \
 -d '{"code": "while(true) {}"}'



