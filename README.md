# Golang-Hash Map Project
Golang Hash Map Project


## Building
1. Download [joeljunstrom's go-luhn library](https://github.com/joeljunstrom/go-luhn)

```
go get -t github.com/HauptJ/go-luhn
```

2. Build the executible

```
go build

```

## Running

1. File input
```
./Golang-Hash-Table input.txt
```

2. STDIN
```
./Golang-Hash_table
<Text>
CRTL+D
```

## Testing

```
go test
```

## Design
1. Why Golang?
- Golang is simple backend programming language that is used in many distributed computing applications such as Kubernetes, Docker, Packer, and Terraform. I like Go's philosophy of keeping the language as simple as possible, as well as the fact that it hs the performance beneifits of a compiled language.

2. Why a Hash Map?
- I noticed that the goal of this project was to update and to keep track of the balances of specified card holders by their names. In order to optimize lookup time, I felt a hash map would serve this purpose best because of its constant O(1), lookup time. For the keys, I decided to use the cardholders name, and for the values, I decided to use an object that kept track of the card's spending limit, current balance, as well as its validity.

3. Why joeljunstrom's go-luhn library?
- I noticed he had created a well documented and tested Luhn 10 library in Golang and instead of having to re-invent the wheel, I decided to use his library.