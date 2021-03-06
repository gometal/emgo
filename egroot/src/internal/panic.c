#include <internal/types.h>
#include <internal.h>

void panic(interface i) {
	for (;;) {
		if (internal$Panic != nil) {
			internal$Panic(i);
		}
	}
}

void panicIC() {
	panic(INTERFACE(EGSTL("interface conversion"), &string$$));
}

void panicIndex() {
	panic(INTERFACE(EGSTL("index out of range"), &string$$));
}
