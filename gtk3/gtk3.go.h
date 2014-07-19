#ifndef GO_GTK3_H
#define GO_GTK3_H

#include <gtk/gtk.h>
#include <stdlib.h>

static inline gchar* toGstr(const char* s) { return (gchar*)s; }

static inline GtkWindow* toGWindow(GtkWidget* w) { return GTK_WINDOW(w); }

#endif
