/*
Package rossby implements a distributed, encrypted message service that is intended to
be geo-replicated. Rossby's goal is to facilitate the sending and receiving of messages
while also ensuring that as much work as possible is pushed to the client side.
Requiring clients to handle most of the encryption and message management ensures that
rossby has as little detail as possible to be exposed to any security vulnerabilities.
*/
package rossby
