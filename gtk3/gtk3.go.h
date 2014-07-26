#ifndef GO_GTK3_H
#define GO_GTK3_H

#include <gtk/gtk.h>
#include <stdlib.h>

static inline gchar* toGstr(const char* s) { return (gchar*)s; }

static inline GtkContainer* toGContainer(GtkWidget* w) { return GTK_CONTAINER(w); }
static inline GtkWindow* toGWindow(GtkWidget* w) { return GTK_WINDOW(w); }
static inline GtkWidget* toGWidget(void* w) { return GTK_WIDGET(w); }

#endif
