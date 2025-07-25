#include <QAction>
#include <QActionEvent>
#include <QByteArray>
#include <QChildEvent>
#include <QCloseEvent>
#include <QContextMenuEvent>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEvent>
#include <QFocusEvent>
#include <QFont>
#include <QHideEvent>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QList>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEngine>
#include <QPaintEvent>
#include <QPainter>
#include <QPoint>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTabletEvent>
#include <QTermWidget>
#include <QTextCodec>
#include <QTimerEvent>
#include <QUrl>
#include <QVariant>
#include <QWheelEvent>
#include <QWidget>
#include <qtermwidget.h>
#include "gen_qtermwidget.h"

#ifdef __cplusplus
extern "C" {
#endif

void miqt_exec_callback_QTermWidget_finished(intptr_t);
void miqt_exec_callback_QTermWidget_copyAvailable(intptr_t, bool);
void miqt_exec_callback_QTermWidget_termGetFocus(intptr_t);
void miqt_exec_callback_QTermWidget_termLostFocus(intptr_t);
void miqt_exec_callback_QTermWidget_termKeyPressed(intptr_t, QKeyEvent*);
void miqt_exec_callback_QTermWidget_urlActivated(intptr_t, QUrl*, bool);
void miqt_exec_callback_QTermWidget_bell(intptr_t, struct miqt_string);
void miqt_exec_callback_QTermWidget_activity(intptr_t);
void miqt_exec_callback_QTermWidget_silence(intptr_t);
void miqt_exec_callback_QTermWidget_sendData(intptr_t, const char*, int);
void miqt_exec_callback_QTermWidget_profileChanged(intptr_t, struct miqt_string);
void miqt_exec_callback_QTermWidget_titleChanged(intptr_t);
void miqt_exec_callback_QTermWidget_rightMouseButtonPressed(intptr_t);
void miqt_exec_callback_QTermWidget_receivedData(intptr_t, struct miqt_string);
void miqt_exec_callback_QTermWidget_resizeEvent(QTermWidget*, intptr_t, QResizeEvent*);
void miqt_exec_callback_QTermWidget_keyPressEvent(QTermWidget*, intptr_t, QKeyEvent*);
int miqt_exec_callback_QTermWidget_devType(const QTermWidget*, intptr_t);
void miqt_exec_callback_QTermWidget_setVisible(QTermWidget*, intptr_t, bool);
QSize* miqt_exec_callback_QTermWidget_sizeHint(const QTermWidget*, intptr_t);
QSize* miqt_exec_callback_QTermWidget_minimumSizeHint(const QTermWidget*, intptr_t);
int miqt_exec_callback_QTermWidget_heightForWidth(const QTermWidget*, intptr_t, int);
bool miqt_exec_callback_QTermWidget_hasHeightForWidth(const QTermWidget*, intptr_t);
QPaintEngine* miqt_exec_callback_QTermWidget_paintEngine(const QTermWidget*, intptr_t);
bool miqt_exec_callback_QTermWidget_event(QTermWidget*, intptr_t, QEvent*);
void miqt_exec_callback_QTermWidget_mousePressEvent(QTermWidget*, intptr_t, QMouseEvent*);
void miqt_exec_callback_QTermWidget_mouseReleaseEvent(QTermWidget*, intptr_t, QMouseEvent*);
void miqt_exec_callback_QTermWidget_mouseDoubleClickEvent(QTermWidget*, intptr_t, QMouseEvent*);
void miqt_exec_callback_QTermWidget_mouseMoveEvent(QTermWidget*, intptr_t, QMouseEvent*);
void miqt_exec_callback_QTermWidget_wheelEvent(QTermWidget*, intptr_t, QWheelEvent*);
void miqt_exec_callback_QTermWidget_keyReleaseEvent(QTermWidget*, intptr_t, QKeyEvent*);
void miqt_exec_callback_QTermWidget_focusInEvent(QTermWidget*, intptr_t, QFocusEvent*);
void miqt_exec_callback_QTermWidget_focusOutEvent(QTermWidget*, intptr_t, QFocusEvent*);
void miqt_exec_callback_QTermWidget_enterEvent(QTermWidget*, intptr_t, QEvent*);
void miqt_exec_callback_QTermWidget_leaveEvent(QTermWidget*, intptr_t, QEvent*);
void miqt_exec_callback_QTermWidget_paintEvent(QTermWidget*, intptr_t, QPaintEvent*);
void miqt_exec_callback_QTermWidget_moveEvent(QTermWidget*, intptr_t, QMoveEvent*);
void miqt_exec_callback_QTermWidget_closeEvent(QTermWidget*, intptr_t, QCloseEvent*);
void miqt_exec_callback_QTermWidget_contextMenuEvent(QTermWidget*, intptr_t, QContextMenuEvent*);
void miqt_exec_callback_QTermWidget_tabletEvent(QTermWidget*, intptr_t, QTabletEvent*);
void miqt_exec_callback_QTermWidget_actionEvent(QTermWidget*, intptr_t, QActionEvent*);
void miqt_exec_callback_QTermWidget_dragEnterEvent(QTermWidget*, intptr_t, QDragEnterEvent*);
void miqt_exec_callback_QTermWidget_dragMoveEvent(QTermWidget*, intptr_t, QDragMoveEvent*);
void miqt_exec_callback_QTermWidget_dragLeaveEvent(QTermWidget*, intptr_t, QDragLeaveEvent*);
void miqt_exec_callback_QTermWidget_dropEvent(QTermWidget*, intptr_t, QDropEvent*);
void miqt_exec_callback_QTermWidget_showEvent(QTermWidget*, intptr_t, QShowEvent*);
void miqt_exec_callback_QTermWidget_hideEvent(QTermWidget*, intptr_t, QHideEvent*);
bool miqt_exec_callback_QTermWidget_nativeEvent(QTermWidget*, intptr_t, struct miqt_string, void*, long*);
void miqt_exec_callback_QTermWidget_changeEvent(QTermWidget*, intptr_t, QEvent*);
int miqt_exec_callback_QTermWidget_metric(const QTermWidget*, intptr_t, int);
void miqt_exec_callback_QTermWidget_initPainter(const QTermWidget*, intptr_t, QPainter*);
QPaintDevice* miqt_exec_callback_QTermWidget_redirected(const QTermWidget*, intptr_t, QPoint*);
QPainter* miqt_exec_callback_QTermWidget_sharedPainter(const QTermWidget*, intptr_t);
void miqt_exec_callback_QTermWidget_inputMethodEvent(QTermWidget*, intptr_t, QInputMethodEvent*);
QVariant* miqt_exec_callback_QTermWidget_inputMethodQuery(const QTermWidget*, intptr_t, int);
bool miqt_exec_callback_QTermWidget_focusNextPrevChild(QTermWidget*, intptr_t, bool);
bool miqt_exec_callback_QTermWidget_eventFilter(QTermWidget*, intptr_t, QObject*, QEvent*);
void miqt_exec_callback_QTermWidget_timerEvent(QTermWidget*, intptr_t, QTimerEvent*);
void miqt_exec_callback_QTermWidget_childEvent(QTermWidget*, intptr_t, QChildEvent*);
void miqt_exec_callback_QTermWidget_customEvent(QTermWidget*, intptr_t, QEvent*);
void miqt_exec_callback_QTermWidget_connectNotify(QTermWidget*, intptr_t, QMetaMethod*);
void miqt_exec_callback_QTermWidget_disconnectNotify(QTermWidget*, intptr_t, QMetaMethod*);
#ifdef __cplusplus
} /* extern C */
#endif

class MiqtVirtualQTermWidget final : public QTermWidget {
public:

	MiqtVirtualQTermWidget(QWidget* parent): QTermWidget(parent) {};
	MiqtVirtualQTermWidget(int startnow): QTermWidget(startnow) {};
	MiqtVirtualQTermWidget(): QTermWidget() {};
	MiqtVirtualQTermWidget(int startnow, QWidget* parent): QTermWidget(startnow, parent) {};

	virtual ~MiqtVirtualQTermWidget() override = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__resizeEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void resizeEvent(QResizeEvent* param1) override {
		if (handle__resizeEvent == 0) {
			QTermWidget::resizeEvent(param1);
			return;
		}
		
		QResizeEvent* sigval1 = param1;

		miqt_exec_callback_QTermWidget_resizeEvent(this, handle__resizeEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_resizeEvent(void* self, QResizeEvent* param1);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__keyPressEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void keyPressEvent(QKeyEvent* event) override {
		if (handle__keyPressEvent == 0) {
			QTermWidget::keyPressEvent(event);
			return;
		}
		
		QKeyEvent* sigval1 = event;

		miqt_exec_callback_QTermWidget_keyPressEvent(this, handle__keyPressEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_keyPressEvent(void* self, QKeyEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__devType = 0;

	// Subclass to allow providing a Go implementation
	virtual int devType() const override {
		if (handle__devType == 0) {
			return QTermWidget::devType();
		}
		

		int callback_return_value = miqt_exec_callback_QTermWidget_devType(this, handle__devType);

		return static_cast<int>(callback_return_value);
	}

	friend int QTermWidget_virtualbase_devType(const void* self);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__setVisible = 0;

	// Subclass to allow providing a Go implementation
	virtual void setVisible(bool visible) override {
		if (handle__setVisible == 0) {
			QTermWidget::setVisible(visible);
			return;
		}
		
		bool sigval1 = visible;

		miqt_exec_callback_QTermWidget_setVisible(this, handle__setVisible, sigval1);

		
	}

	friend void QTermWidget_virtualbase_setVisible(void* self, bool visible);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__sizeHint = 0;

	// Subclass to allow providing a Go implementation
	virtual QSize sizeHint() const override {
		if (handle__sizeHint == 0) {
			return QTermWidget::sizeHint();
		}
		

		QSize* callback_return_value = miqt_exec_callback_QTermWidget_sizeHint(this, handle__sizeHint);

		return *callback_return_value;
	}

	friend QSize* QTermWidget_virtualbase_sizeHint(const void* self);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__minimumSizeHint = 0;

	// Subclass to allow providing a Go implementation
	virtual QSize minimumSizeHint() const override {
		if (handle__minimumSizeHint == 0) {
			return QTermWidget::minimumSizeHint();
		}
		

		QSize* callback_return_value = miqt_exec_callback_QTermWidget_minimumSizeHint(this, handle__minimumSizeHint);

		return *callback_return_value;
	}

	friend QSize* QTermWidget_virtualbase_minimumSizeHint(const void* self);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__heightForWidth = 0;

	// Subclass to allow providing a Go implementation
	virtual int heightForWidth(int param1) const override {
		if (handle__heightForWidth == 0) {
			return QTermWidget::heightForWidth(param1);
		}
		
		int sigval1 = param1;

		int callback_return_value = miqt_exec_callback_QTermWidget_heightForWidth(this, handle__heightForWidth, sigval1);

		return static_cast<int>(callback_return_value);
	}

	friend int QTermWidget_virtualbase_heightForWidth(const void* self, int param1);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__hasHeightForWidth = 0;

	// Subclass to allow providing a Go implementation
	virtual bool hasHeightForWidth() const override {
		if (handle__hasHeightForWidth == 0) {
			return QTermWidget::hasHeightForWidth();
		}
		

		bool callback_return_value = miqt_exec_callback_QTermWidget_hasHeightForWidth(this, handle__hasHeightForWidth);

		return callback_return_value;
	}

	friend bool QTermWidget_virtualbase_hasHeightForWidth(const void* self);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__paintEngine = 0;

	// Subclass to allow providing a Go implementation
	virtual QPaintEngine* paintEngine() const override {
		if (handle__paintEngine == 0) {
			return QTermWidget::paintEngine();
		}
		

		QPaintEngine* callback_return_value = miqt_exec_callback_QTermWidget_paintEngine(this, handle__paintEngine);

		return callback_return_value;
	}

	friend QPaintEngine* QTermWidget_virtualbase_paintEngine(const void* self);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__event = 0;

	// Subclass to allow providing a Go implementation
	virtual bool event(QEvent* event) override {
		if (handle__event == 0) {
			return QTermWidget::event(event);
		}
		
		QEvent* sigval1 = event;

		bool callback_return_value = miqt_exec_callback_QTermWidget_event(this, handle__event, sigval1);

		return callback_return_value;
	}

	friend bool QTermWidget_virtualbase_event(void* self, QEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__mousePressEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void mousePressEvent(QMouseEvent* event) override {
		if (handle__mousePressEvent == 0) {
			QTermWidget::mousePressEvent(event);
			return;
		}
		
		QMouseEvent* sigval1 = event;

		miqt_exec_callback_QTermWidget_mousePressEvent(this, handle__mousePressEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_mousePressEvent(void* self, QMouseEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__mouseReleaseEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void mouseReleaseEvent(QMouseEvent* event) override {
		if (handle__mouseReleaseEvent == 0) {
			QTermWidget::mouseReleaseEvent(event);
			return;
		}
		
		QMouseEvent* sigval1 = event;

		miqt_exec_callback_QTermWidget_mouseReleaseEvent(this, handle__mouseReleaseEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_mouseReleaseEvent(void* self, QMouseEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__mouseDoubleClickEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void mouseDoubleClickEvent(QMouseEvent* event) override {
		if (handle__mouseDoubleClickEvent == 0) {
			QTermWidget::mouseDoubleClickEvent(event);
			return;
		}
		
		QMouseEvent* sigval1 = event;

		miqt_exec_callback_QTermWidget_mouseDoubleClickEvent(this, handle__mouseDoubleClickEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_mouseDoubleClickEvent(void* self, QMouseEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__mouseMoveEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void mouseMoveEvent(QMouseEvent* event) override {
		if (handle__mouseMoveEvent == 0) {
			QTermWidget::mouseMoveEvent(event);
			return;
		}
		
		QMouseEvent* sigval1 = event;

		miqt_exec_callback_QTermWidget_mouseMoveEvent(this, handle__mouseMoveEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_mouseMoveEvent(void* self, QMouseEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__wheelEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void wheelEvent(QWheelEvent* event) override {
		if (handle__wheelEvent == 0) {
			QTermWidget::wheelEvent(event);
			return;
		}
		
		QWheelEvent* sigval1 = event;

		miqt_exec_callback_QTermWidget_wheelEvent(this, handle__wheelEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_wheelEvent(void* self, QWheelEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__keyReleaseEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void keyReleaseEvent(QKeyEvent* event) override {
		if (handle__keyReleaseEvent == 0) {
			QTermWidget::keyReleaseEvent(event);
			return;
		}
		
		QKeyEvent* sigval1 = event;

		miqt_exec_callback_QTermWidget_keyReleaseEvent(this, handle__keyReleaseEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_keyReleaseEvent(void* self, QKeyEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__focusInEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void focusInEvent(QFocusEvent* event) override {
		if (handle__focusInEvent == 0) {
			QTermWidget::focusInEvent(event);
			return;
		}
		
		QFocusEvent* sigval1 = event;

		miqt_exec_callback_QTermWidget_focusInEvent(this, handle__focusInEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_focusInEvent(void* self, QFocusEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__focusOutEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void focusOutEvent(QFocusEvent* event) override {
		if (handle__focusOutEvent == 0) {
			QTermWidget::focusOutEvent(event);
			return;
		}
		
		QFocusEvent* sigval1 = event;

		miqt_exec_callback_QTermWidget_focusOutEvent(this, handle__focusOutEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_focusOutEvent(void* self, QFocusEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__enterEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void enterEvent(QEvent* event) override {
		if (handle__enterEvent == 0) {
			QTermWidget::enterEvent(event);
			return;
		}
		
		QEvent* sigval1 = event;

		miqt_exec_callback_QTermWidget_enterEvent(this, handle__enterEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_enterEvent(void* self, QEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__leaveEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void leaveEvent(QEvent* event) override {
		if (handle__leaveEvent == 0) {
			QTermWidget::leaveEvent(event);
			return;
		}
		
		QEvent* sigval1 = event;

		miqt_exec_callback_QTermWidget_leaveEvent(this, handle__leaveEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_leaveEvent(void* self, QEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__paintEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void paintEvent(QPaintEvent* event) override {
		if (handle__paintEvent == 0) {
			QTermWidget::paintEvent(event);
			return;
		}
		
		QPaintEvent* sigval1 = event;

		miqt_exec_callback_QTermWidget_paintEvent(this, handle__paintEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_paintEvent(void* self, QPaintEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__moveEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void moveEvent(QMoveEvent* event) override {
		if (handle__moveEvent == 0) {
			QTermWidget::moveEvent(event);
			return;
		}
		
		QMoveEvent* sigval1 = event;

		miqt_exec_callback_QTermWidget_moveEvent(this, handle__moveEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_moveEvent(void* self, QMoveEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__closeEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void closeEvent(QCloseEvent* event) override {
		if (handle__closeEvent == 0) {
			QTermWidget::closeEvent(event);
			return;
		}
		
		QCloseEvent* sigval1 = event;

		miqt_exec_callback_QTermWidget_closeEvent(this, handle__closeEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_closeEvent(void* self, QCloseEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__contextMenuEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void contextMenuEvent(QContextMenuEvent* event) override {
		if (handle__contextMenuEvent == 0) {
			QTermWidget::contextMenuEvent(event);
			return;
		}
		
		QContextMenuEvent* sigval1 = event;

		miqt_exec_callback_QTermWidget_contextMenuEvent(this, handle__contextMenuEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_contextMenuEvent(void* self, QContextMenuEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__tabletEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void tabletEvent(QTabletEvent* event) override {
		if (handle__tabletEvent == 0) {
			QTermWidget::tabletEvent(event);
			return;
		}
		
		QTabletEvent* sigval1 = event;

		miqt_exec_callback_QTermWidget_tabletEvent(this, handle__tabletEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_tabletEvent(void* self, QTabletEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__actionEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void actionEvent(QActionEvent* event) override {
		if (handle__actionEvent == 0) {
			QTermWidget::actionEvent(event);
			return;
		}
		
		QActionEvent* sigval1 = event;

		miqt_exec_callback_QTermWidget_actionEvent(this, handle__actionEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_actionEvent(void* self, QActionEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__dragEnterEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void dragEnterEvent(QDragEnterEvent* event) override {
		if (handle__dragEnterEvent == 0) {
			QTermWidget::dragEnterEvent(event);
			return;
		}
		
		QDragEnterEvent* sigval1 = event;

		miqt_exec_callback_QTermWidget_dragEnterEvent(this, handle__dragEnterEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_dragEnterEvent(void* self, QDragEnterEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__dragMoveEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void dragMoveEvent(QDragMoveEvent* event) override {
		if (handle__dragMoveEvent == 0) {
			QTermWidget::dragMoveEvent(event);
			return;
		}
		
		QDragMoveEvent* sigval1 = event;

		miqt_exec_callback_QTermWidget_dragMoveEvent(this, handle__dragMoveEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_dragMoveEvent(void* self, QDragMoveEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__dragLeaveEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void dragLeaveEvent(QDragLeaveEvent* event) override {
		if (handle__dragLeaveEvent == 0) {
			QTermWidget::dragLeaveEvent(event);
			return;
		}
		
		QDragLeaveEvent* sigval1 = event;

		miqt_exec_callback_QTermWidget_dragLeaveEvent(this, handle__dragLeaveEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_dragLeaveEvent(void* self, QDragLeaveEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__dropEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void dropEvent(QDropEvent* event) override {
		if (handle__dropEvent == 0) {
			QTermWidget::dropEvent(event);
			return;
		}
		
		QDropEvent* sigval1 = event;

		miqt_exec_callback_QTermWidget_dropEvent(this, handle__dropEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_dropEvent(void* self, QDropEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__showEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void showEvent(QShowEvent* event) override {
		if (handle__showEvent == 0) {
			QTermWidget::showEvent(event);
			return;
		}
		
		QShowEvent* sigval1 = event;

		miqt_exec_callback_QTermWidget_showEvent(this, handle__showEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_showEvent(void* self, QShowEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__hideEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void hideEvent(QHideEvent* event) override {
		if (handle__hideEvent == 0) {
			QTermWidget::hideEvent(event);
			return;
		}
		
		QHideEvent* sigval1 = event;

		miqt_exec_callback_QTermWidget_hideEvent(this, handle__hideEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_hideEvent(void* self, QHideEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__nativeEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual bool nativeEvent(const QByteArray& eventType, void* message, long* result) override {
		if (handle__nativeEvent == 0) {
			return QTermWidget::nativeEvent(eventType, message, result);
		}
		
		const QByteArray eventType_qb = eventType;
		struct miqt_string eventType_ms;
		eventType_ms.len = eventType_qb.length();
		eventType_ms.data = static_cast<char*>(malloc(eventType_ms.len));
		memcpy(eventType_ms.data, eventType_qb.data(), eventType_ms.len);
		struct miqt_string sigval1 = eventType_ms;
		void* sigval2 = message;
		long* sigval3 = result;

		bool callback_return_value = miqt_exec_callback_QTermWidget_nativeEvent(this, handle__nativeEvent, sigval1, sigval2, sigval3);

		return callback_return_value;
	}

	friend bool QTermWidget_virtualbase_nativeEvent(void* self, struct miqt_string eventType, void* message, long* result);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__changeEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void changeEvent(QEvent* param1) override {
		if (handle__changeEvent == 0) {
			QTermWidget::changeEvent(param1);
			return;
		}
		
		QEvent* sigval1 = param1;

		miqt_exec_callback_QTermWidget_changeEvent(this, handle__changeEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_changeEvent(void* self, QEvent* param1);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__metric = 0;

	// Subclass to allow providing a Go implementation
	virtual int metric(QPaintDevice::PaintDeviceMetric param1) const override {
		if (handle__metric == 0) {
			return QTermWidget::metric(param1);
		}
		
		QPaintDevice::PaintDeviceMetric param1_ret = param1;
		int sigval1 = static_cast<int>(param1_ret);

		int callback_return_value = miqt_exec_callback_QTermWidget_metric(this, handle__metric, sigval1);

		return static_cast<int>(callback_return_value);
	}

	friend int QTermWidget_virtualbase_metric(const void* self, int param1);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__initPainter = 0;

	// Subclass to allow providing a Go implementation
	virtual void initPainter(QPainter* painter) const override {
		if (handle__initPainter == 0) {
			QTermWidget::initPainter(painter);
			return;
		}
		
		QPainter* sigval1 = painter;

		miqt_exec_callback_QTermWidget_initPainter(this, handle__initPainter, sigval1);

		
	}

	friend void QTermWidget_virtualbase_initPainter(const void* self, QPainter* painter);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__redirected = 0;

	// Subclass to allow providing a Go implementation
	virtual QPaintDevice* redirected(QPoint* offset) const override {
		if (handle__redirected == 0) {
			return QTermWidget::redirected(offset);
		}
		
		QPoint* sigval1 = offset;

		QPaintDevice* callback_return_value = miqt_exec_callback_QTermWidget_redirected(this, handle__redirected, sigval1);

		return callback_return_value;
	}

	friend QPaintDevice* QTermWidget_virtualbase_redirected(const void* self, QPoint* offset);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__sharedPainter = 0;

	// Subclass to allow providing a Go implementation
	virtual QPainter* sharedPainter() const override {
		if (handle__sharedPainter == 0) {
			return QTermWidget::sharedPainter();
		}
		

		QPainter* callback_return_value = miqt_exec_callback_QTermWidget_sharedPainter(this, handle__sharedPainter);

		return callback_return_value;
	}

	friend QPainter* QTermWidget_virtualbase_sharedPainter(const void* self);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__inputMethodEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void inputMethodEvent(QInputMethodEvent* param1) override {
		if (handle__inputMethodEvent == 0) {
			QTermWidget::inputMethodEvent(param1);
			return;
		}
		
		QInputMethodEvent* sigval1 = param1;

		miqt_exec_callback_QTermWidget_inputMethodEvent(this, handle__inputMethodEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_inputMethodEvent(void* self, QInputMethodEvent* param1);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__inputMethodQuery = 0;

	// Subclass to allow providing a Go implementation
	virtual QVariant inputMethodQuery(Qt::InputMethodQuery param1) const override {
		if (handle__inputMethodQuery == 0) {
			return QTermWidget::inputMethodQuery(param1);
		}
		
		Qt::InputMethodQuery param1_ret = param1;
		int sigval1 = static_cast<int>(param1_ret);

		QVariant* callback_return_value = miqt_exec_callback_QTermWidget_inputMethodQuery(this, handle__inputMethodQuery, sigval1);

		return *callback_return_value;
	}

	friend QVariant* QTermWidget_virtualbase_inputMethodQuery(const void* self, int param1);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__focusNextPrevChild = 0;

	// Subclass to allow providing a Go implementation
	virtual bool focusNextPrevChild(bool next) override {
		if (handle__focusNextPrevChild == 0) {
			return QTermWidget::focusNextPrevChild(next);
		}
		
		bool sigval1 = next;

		bool callback_return_value = miqt_exec_callback_QTermWidget_focusNextPrevChild(this, handle__focusNextPrevChild, sigval1);

		return callback_return_value;
	}

	friend bool QTermWidget_virtualbase_focusNextPrevChild(void* self, bool next);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__eventFilter = 0;

	// Subclass to allow providing a Go implementation
	virtual bool eventFilter(QObject* watched, QEvent* event) override {
		if (handle__eventFilter == 0) {
			return QTermWidget::eventFilter(watched, event);
		}
		
		QObject* sigval1 = watched;
		QEvent* sigval2 = event;

		bool callback_return_value = miqt_exec_callback_QTermWidget_eventFilter(this, handle__eventFilter, sigval1, sigval2);

		return callback_return_value;
	}

	friend bool QTermWidget_virtualbase_eventFilter(void* self, QObject* watched, QEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__timerEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void timerEvent(QTimerEvent* event) override {
		if (handle__timerEvent == 0) {
			QTermWidget::timerEvent(event);
			return;
		}
		
		QTimerEvent* sigval1 = event;

		miqt_exec_callback_QTermWidget_timerEvent(this, handle__timerEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_timerEvent(void* self, QTimerEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__childEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void childEvent(QChildEvent* event) override {
		if (handle__childEvent == 0) {
			QTermWidget::childEvent(event);
			return;
		}
		
		QChildEvent* sigval1 = event;

		miqt_exec_callback_QTermWidget_childEvent(this, handle__childEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_childEvent(void* self, QChildEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__customEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void customEvent(QEvent* event) override {
		if (handle__customEvent == 0) {
			QTermWidget::customEvent(event);
			return;
		}
		
		QEvent* sigval1 = event;

		miqt_exec_callback_QTermWidget_customEvent(this, handle__customEvent, sigval1);

		
	}

	friend void QTermWidget_virtualbase_customEvent(void* self, QEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__connectNotify = 0;

	// Subclass to allow providing a Go implementation
	virtual void connectNotify(const QMetaMethod& signal) override {
		if (handle__connectNotify == 0) {
			QTermWidget::connectNotify(signal);
			return;
		}
		
		const QMetaMethod& signal_ret = signal;
		// Cast returned reference into pointer
		QMetaMethod* sigval1 = const_cast<QMetaMethod*>(&signal_ret);

		miqt_exec_callback_QTermWidget_connectNotify(this, handle__connectNotify, sigval1);

		
	}

	friend void QTermWidget_virtualbase_connectNotify(void* self, QMetaMethod* signal);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__disconnectNotify = 0;

	// Subclass to allow providing a Go implementation
	virtual void disconnectNotify(const QMetaMethod& signal) override {
		if (handle__disconnectNotify == 0) {
			QTermWidget::disconnectNotify(signal);
			return;
		}
		
		const QMetaMethod& signal_ret = signal;
		// Cast returned reference into pointer
		QMetaMethod* sigval1 = const_cast<QMetaMethod*>(&signal_ret);

		miqt_exec_callback_QTermWidget_disconnectNotify(this, handle__disconnectNotify, sigval1);

		
	}

	friend void QTermWidget_virtualbase_disconnectNotify(void* self, QMetaMethod* signal);

	// Wrappers to allow calling protected methods:
	friend void QTermWidget_protectedbase_sessionFinished(bool* _dynamic_cast_ok, void* self);
	friend void QTermWidget_protectedbase_selectionChanged(bool* _dynamic_cast_ok, void* self, bool textSelected);
	friend void QTermWidget_protectedbase_updateMicroFocus(bool* _dynamic_cast_ok, void* self);
	friend void QTermWidget_protectedbase_create(bool* _dynamic_cast_ok, void* self);
	friend void QTermWidget_protectedbase_destroy(bool* _dynamic_cast_ok, void* self);
	friend bool QTermWidget_protectedbase_focusNextChild(bool* _dynamic_cast_ok, void* self);
	friend bool QTermWidget_protectedbase_focusPreviousChild(bool* _dynamic_cast_ok, void* self);
	friend QObject* QTermWidget_protectedbase_sender(bool* _dynamic_cast_ok, const void* self);
	friend int QTermWidget_protectedbase_senderSignalIndex(bool* _dynamic_cast_ok, const void* self);
	friend int QTermWidget_protectedbase_receivers(bool* _dynamic_cast_ok, const void* self, const char* signal);
	friend bool QTermWidget_protectedbase_isSignalConnected(bool* _dynamic_cast_ok, const void* self, QMetaMethod* signal);
};

QTermWidget* QTermWidget_new(QWidget* parent) {
	return new MiqtVirtualQTermWidget(parent);
}

QTermWidget* QTermWidget_new2(int startnow) {
	return new MiqtVirtualQTermWidget(static_cast<int>(startnow));
}

QTermWidget* QTermWidget_new3() {
	return new MiqtVirtualQTermWidget();
}

QTermWidget* QTermWidget_new4(int startnow, QWidget* parent) {
	return new MiqtVirtualQTermWidget(static_cast<int>(startnow), parent);
}

void QTermWidget_virtbase(QTermWidget* src, QWidget** outptr_QWidget) {
	*outptr_QWidget = static_cast<QWidget*>(src);
}

QMetaObject* QTermWidget_metaObject(const QTermWidget* self) {
	return (QMetaObject*) self->metaObject();
}

void* QTermWidget_metacast(QTermWidget* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QTermWidget_tr(const char* s) {
	QString _ret = QTermWidget::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QTermWidget_trUtf8(const char* s) {
	QString _ret = QTermWidget::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QSize* QTermWidget_sizeHint(const QTermWidget* self) {
	return new QSize(self->sizeHint());
}

void QTermWidget_setTerminalSizeHint(QTermWidget* self, bool on) {
	self->setTerminalSizeHint(on);
}

bool QTermWidget_terminalSizeHint(QTermWidget* self) {
	return self->terminalSizeHint();
}

void QTermWidget_startShellProgram(QTermWidget* self) {
	self->startShellProgram();
}

void QTermWidget_startTerminalTeletype(QTermWidget* self) {
	self->startTerminalTeletype();
}

int QTermWidget_getShellPID(QTermWidget* self) {
	return self->getShellPID();
}

void QTermWidget_changeDir(QTermWidget* self, struct miqt_string dir) {
	QString dir_QString = QString::fromUtf8(dir.data, dir.len);
	self->changeDir(dir_QString);
}

void QTermWidget_setTerminalFont(QTermWidget* self, QFont* font) {
	self->setTerminalFont(*font);
}

QFont* QTermWidget_getTerminalFont(QTermWidget* self) {
	return new QFont(self->getTerminalFont());
}

void QTermWidget_setTerminalOpacity(QTermWidget* self, double level) {
	self->setTerminalOpacity(static_cast<qreal>(level));
}

void QTermWidget_setTerminalBackgroundImage(QTermWidget* self, struct miqt_string backgroundImage) {
	QString backgroundImage_QString = QString::fromUtf8(backgroundImage.data, backgroundImage.len);
	self->setTerminalBackgroundImage(backgroundImage_QString);
}

void QTermWidget_setEnvironment(QTermWidget* self, struct miqt_array /* of struct miqt_string */  environment) {
	QStringList environment_QList;
	environment_QList.reserve(environment.len);
	struct miqt_string* environment_arr = static_cast<struct miqt_string*>(environment.data);
	for(size_t i = 0; i < environment.len; ++i) {
		QString environment_arr_i_QString = QString::fromUtf8(environment_arr[i].data, environment_arr[i].len);
		environment_QList.push_back(environment_arr_i_QString);
	}
	self->setEnvironment(environment_QList);
}

void QTermWidget_setShellProgram(QTermWidget* self, struct miqt_string progname) {
	QString progname_QString = QString::fromUtf8(progname.data, progname.len);
	self->setShellProgram(progname_QString);
}

void QTermWidget_setWorkingDirectory(QTermWidget* self, struct miqt_string dir) {
	QString dir_QString = QString::fromUtf8(dir.data, dir.len);
	self->setWorkingDirectory(dir_QString);
}

struct miqt_string QTermWidget_workingDirectory(QTermWidget* self) {
	QString _ret = self->workingDirectory();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QTermWidget_setArgs(QTermWidget* self, struct miqt_array /* of struct miqt_string */  args) {
	QStringList args_QList;
	args_QList.reserve(args.len);
	struct miqt_string* args_arr = static_cast<struct miqt_string*>(args.data);
	for(size_t i = 0; i < args.len; ++i) {
		QString args_arr_i_QString = QString::fromUtf8(args_arr[i].data, args_arr[i].len);
		args_QList.push_back(args_arr_i_QString);
	}
	self->setArgs(args_QList);
}

void QTermWidget_setTextCodec(QTermWidget* self, QTextCodec* codec) {
	self->setTextCodec(codec);
}

void QTermWidget_setColorScheme(QTermWidget* self, struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	self->setColorScheme(name_QString);
}

struct miqt_array /* of struct miqt_string */  QTermWidget_availableColorSchemes() {
	QStringList _ret = QTermWidget::availableColorSchemes();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_b.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_b.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

void QTermWidget_addCustomColorSchemeDir(struct miqt_string custom_dir) {
	QString custom_dir_QString = QString::fromUtf8(custom_dir.data, custom_dir.len);
	QTermWidget::addCustomColorSchemeDir(custom_dir_QString);
}

void QTermWidget_setHistorySize(QTermWidget* self, int lines) {
	self->setHistorySize(static_cast<int>(lines));
}

void QTermWidget_setScrollBarPosition(QTermWidget* self, int scrollBarPosition) {
	self->setScrollBarPosition(static_cast<QTermWidget::ScrollBarPosition>(scrollBarPosition));
}

void QTermWidget_scrollToEnd(QTermWidget* self) {
	self->scrollToEnd();
}

void QTermWidget_sendText(QTermWidget* self, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	self->sendText(text_QString);
}

void QTermWidget_setFlowControlEnabled(QTermWidget* self, bool enabled) {
	self->setFlowControlEnabled(enabled);
}

bool QTermWidget_flowControlEnabled(QTermWidget* self) {
	return self->flowControlEnabled();
}

void QTermWidget_setFlowControlWarningEnabled(QTermWidget* self, bool enabled) {
	self->setFlowControlWarningEnabled(enabled);
}

struct miqt_array /* of struct miqt_string */  QTermWidget_availableKeyBindings() {
	QStringList _ret = QTermWidget::availableKeyBindings();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_b.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_b.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_string QTermWidget_keyBindings(QTermWidget* self) {
	QString _ret = self->keyBindings();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QTermWidget_setMotionAfterPasting(QTermWidget* self, int motionAfterPasting) {
	self->setMotionAfterPasting(static_cast<int>(motionAfterPasting));
}

int QTermWidget_historyLinesCount(QTermWidget* self) {
	return self->historyLinesCount();
}

int QTermWidget_screenColumnsCount(QTermWidget* self) {
	return self->screenColumnsCount();
}

int QTermWidget_screenLinesCount(QTermWidget* self) {
	return self->screenLinesCount();
}

void QTermWidget_setSelectionStart(QTermWidget* self, int row, int column) {
	self->setSelectionStart(static_cast<int>(row), static_cast<int>(column));
}

void QTermWidget_setSelectionEnd(QTermWidget* self, int row, int column) {
	self->setSelectionEnd(static_cast<int>(row), static_cast<int>(column));
}

void QTermWidget_getSelectionStart(QTermWidget* self, int* row, int* column) {
	self->getSelectionStart(static_cast<int&>(*row), static_cast<int&>(*column));
}

void QTermWidget_getSelectionEnd(QTermWidget* self, int* row, int* column) {
	self->getSelectionEnd(static_cast<int&>(*row), static_cast<int&>(*column));
}

struct miqt_string QTermWidget_selectedText(QTermWidget* self) {
	QString _ret = self->selectedText();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QTermWidget_setMonitorActivity(QTermWidget* self, bool monitorActivity) {
	self->setMonitorActivity(monitorActivity);
}

void QTermWidget_setMonitorSilence(QTermWidget* self, bool monitorSilence) {
	self->setMonitorSilence(monitorSilence);
}

void QTermWidget_setSilenceTimeout(QTermWidget* self, int seconds) {
	self->setSilenceTimeout(static_cast<int>(seconds));
}

struct miqt_array /* of QAction* */  QTermWidget_filterActions(QTermWidget* self, QPoint* position) {
	QList<QAction *> _ret = self->filterActions(*position);
	// Convert QList<> from C++ memory to manually-managed C memory
	QAction** _arr = static_cast<QAction**>(malloc(sizeof(QAction*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

int QTermWidget_getPtySlaveFd(const QTermWidget* self) {
	return self->getPtySlaveFd();
}

void QTermWidget_setKeyboardCursorShape(QTermWidget* self, int shape) {
	self->setKeyboardCursorShape(static_cast<QTermWidget::KeyboardCursorShape>(shape));
}

void QTermWidget_setBlinkingCursor(QTermWidget* self, bool blink) {
	self->setBlinkingCursor(blink);
}

void QTermWidget_setBidiEnabled(QTermWidget* self, bool enabled) {
	self->setBidiEnabled(enabled);
}

bool QTermWidget_isBidiEnabled(QTermWidget* self) {
	return self->isBidiEnabled();
}

void QTermWidget_setAutoClose(QTermWidget* self, bool autoClose) {
	self->setAutoClose(autoClose);
}

struct miqt_string QTermWidget_title(const QTermWidget* self) {
	QString _ret = self->title();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QTermWidget_icon(const QTermWidget* self) {
	QString _ret = self->icon();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QTermWidget_isTitleChanged(const QTermWidget* self) {
	return self->isTitleChanged();
}

void QTermWidget_bracketText(QTermWidget* self, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	self->bracketText(text_QString);
}

void QTermWidget_finished(QTermWidget* self) {
	self->finished();
}

void QTermWidget_connect_finished(QTermWidget* self, intptr_t slot) {
	MiqtVirtualQTermWidget::connect(self, static_cast<void (QTermWidget::*)()>(&QTermWidget::finished), self, [=]() {
		miqt_exec_callback_QTermWidget_finished(slot);
	});
}

void QTermWidget_copyAvailable(QTermWidget* self, bool param1) {
	self->copyAvailable(param1);
}

void QTermWidget_connect_copyAvailable(QTermWidget* self, intptr_t slot) {
	MiqtVirtualQTermWidget::connect(self, static_cast<void (QTermWidget::*)(bool)>(&QTermWidget::copyAvailable), self, [=](bool param1) {
		bool sigval1 = param1;
		miqt_exec_callback_QTermWidget_copyAvailable(slot, sigval1);
	});
}

void QTermWidget_termGetFocus(QTermWidget* self) {
	self->termGetFocus();
}

void QTermWidget_connect_termGetFocus(QTermWidget* self, intptr_t slot) {
	MiqtVirtualQTermWidget::connect(self, static_cast<void (QTermWidget::*)()>(&QTermWidget::termGetFocus), self, [=]() {
		miqt_exec_callback_QTermWidget_termGetFocus(slot);
	});
}

void QTermWidget_termLostFocus(QTermWidget* self) {
	self->termLostFocus();
}

void QTermWidget_connect_termLostFocus(QTermWidget* self, intptr_t slot) {
	MiqtVirtualQTermWidget::connect(self, static_cast<void (QTermWidget::*)()>(&QTermWidget::termLostFocus), self, [=]() {
		miqt_exec_callback_QTermWidget_termLostFocus(slot);
	});
}

void QTermWidget_termKeyPressed(QTermWidget* self, QKeyEvent* param1) {
	self->termKeyPressed(param1);
}

void QTermWidget_connect_termKeyPressed(QTermWidget* self, intptr_t slot) {
	MiqtVirtualQTermWidget::connect(self, static_cast<void (QTermWidget::*)(QKeyEvent*)>(&QTermWidget::termKeyPressed), self, [=](QKeyEvent* param1) {
		QKeyEvent* sigval1 = param1;
		miqt_exec_callback_QTermWidget_termKeyPressed(slot, sigval1);
	});
}

void QTermWidget_urlActivated(QTermWidget* self, QUrl* param1, bool fromContextMenu) {
	self->urlActivated(*param1, fromContextMenu);
}

void QTermWidget_connect_urlActivated(QTermWidget* self, intptr_t slot) {
	MiqtVirtualQTermWidget::connect(self, static_cast<void (QTermWidget::*)(const QUrl&, bool)>(&QTermWidget::urlActivated), self, [=](const QUrl& param1, bool fromContextMenu) {
		const QUrl& param1_ret = param1;
		// Cast returned reference into pointer
		QUrl* sigval1 = const_cast<QUrl*>(&param1_ret);
		bool sigval2 = fromContextMenu;
		miqt_exec_callback_QTermWidget_urlActivated(slot, sigval1, sigval2);
	});
}

void QTermWidget_bell(QTermWidget* self, struct miqt_string message) {
	QString message_QString = QString::fromUtf8(message.data, message.len);
	self->bell(message_QString);
}

void QTermWidget_connect_bell(QTermWidget* self, intptr_t slot) {
	MiqtVirtualQTermWidget::connect(self, static_cast<void (QTermWidget::*)(const QString&)>(&QTermWidget::bell), self, [=](const QString& message) {
		const QString message_ret = message;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray message_b = message_ret.toUtf8();
		struct miqt_string message_ms;
		message_ms.len = message_b.length();
		message_ms.data = static_cast<char*>(malloc(message_ms.len));
		memcpy(message_ms.data, message_b.data(), message_ms.len);
		struct miqt_string sigval1 = message_ms;
		miqt_exec_callback_QTermWidget_bell(slot, sigval1);
	});
}

void QTermWidget_activity(QTermWidget* self) {
	self->activity();
}

void QTermWidget_connect_activity(QTermWidget* self, intptr_t slot) {
	MiqtVirtualQTermWidget::connect(self, static_cast<void (QTermWidget::*)()>(&QTermWidget::activity), self, [=]() {
		miqt_exec_callback_QTermWidget_activity(slot);
	});
}

void QTermWidget_silence(QTermWidget* self) {
	self->silence();
}

void QTermWidget_connect_silence(QTermWidget* self, intptr_t slot) {
	MiqtVirtualQTermWidget::connect(self, static_cast<void (QTermWidget::*)()>(&QTermWidget::silence), self, [=]() {
		miqt_exec_callback_QTermWidget_silence(slot);
	});
}

void QTermWidget_sendData(QTermWidget* self, const char* param1, int param2) {
	self->sendData(param1, static_cast<int>(param2));
}

void QTermWidget_connect_sendData(QTermWidget* self, intptr_t slot) {
	MiqtVirtualQTermWidget::connect(self, static_cast<void (QTermWidget::*)(const char*, int)>(&QTermWidget::sendData), self, [=](const char* param1, int param2) {
		const char* sigval1 = (const char*) param1;
		int sigval2 = param2;
		miqt_exec_callback_QTermWidget_sendData(slot, sigval1, sigval2);
	});
}

void QTermWidget_profileChanged(QTermWidget* self, struct miqt_string profile) {
	QString profile_QString = QString::fromUtf8(profile.data, profile.len);
	self->profileChanged(profile_QString);
}

void QTermWidget_connect_profileChanged(QTermWidget* self, intptr_t slot) {
	MiqtVirtualQTermWidget::connect(self, static_cast<void (QTermWidget::*)(const QString&)>(&QTermWidget::profileChanged), self, [=](const QString& profile) {
		const QString profile_ret = profile;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray profile_b = profile_ret.toUtf8();
		struct miqt_string profile_ms;
		profile_ms.len = profile_b.length();
		profile_ms.data = static_cast<char*>(malloc(profile_ms.len));
		memcpy(profile_ms.data, profile_b.data(), profile_ms.len);
		struct miqt_string sigval1 = profile_ms;
		miqt_exec_callback_QTermWidget_profileChanged(slot, sigval1);
	});
}

void QTermWidget_titleChanged(QTermWidget* self) {
	self->titleChanged();
}

void QTermWidget_connect_titleChanged(QTermWidget* self, intptr_t slot) {
	MiqtVirtualQTermWidget::connect(self, static_cast<void (QTermWidget::*)()>(&QTermWidget::titleChanged), self, [=]() {
		miqt_exec_callback_QTermWidget_titleChanged(slot);
	});
}

void QTermWidget_rightMouseButtonPressed(QTermWidget* self) {
	self->rightMouseButtonPressed();
}

void QTermWidget_connect_rightMouseButtonPressed(QTermWidget* self, intptr_t slot) {
	MiqtVirtualQTermWidget::connect(self, static_cast<void (QTermWidget::*)()>(&QTermWidget::rightMouseButtonPressed), self, [=]() {
		miqt_exec_callback_QTermWidget_rightMouseButtonPressed(slot);
	});
}

void QTermWidget_receivedData(QTermWidget* self, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	self->receivedData(text_QString);
}

void QTermWidget_connect_receivedData(QTermWidget* self, intptr_t slot) {
	MiqtVirtualQTermWidget::connect(self, static_cast<void (QTermWidget::*)(const QString&)>(&QTermWidget::receivedData), self, [=](const QString& text) {
		const QString text_ret = text;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray text_b = text_ret.toUtf8();
		struct miqt_string text_ms;
		text_ms.len = text_b.length();
		text_ms.data = static_cast<char*>(malloc(text_ms.len));
		memcpy(text_ms.data, text_b.data(), text_ms.len);
		struct miqt_string sigval1 = text_ms;
		miqt_exec_callback_QTermWidget_receivedData(slot, sigval1);
	});
}

void QTermWidget_receivedRemoteData(QTermWidget* self, struct miqt_string data) {
	QByteArray data_QByteArray(data.data, data.len);
	self->receivedRemoteData(data_QByteArray);
}

void QTermWidget_copyClipboard(QTermWidget* self) {
	self->copyClipboard();
}

void QTermWidget_pasteClipboard(QTermWidget* self) {
	self->pasteClipboard();
}

void QTermWidget_pasteSelection(QTermWidget* self) {
	self->pasteSelection();
}

void QTermWidget_zoomIn(QTermWidget* self) {
	self->zoomIn();
}

void QTermWidget_zoomOut(QTermWidget* self) {
	self->zoomOut();
}

void QTermWidget_setSize(QTermWidget* self, QSize* size) {
	self->setSize(*size);
}

void QTermWidget_setKeyBindings(QTermWidget* self, struct miqt_string kb) {
	QString kb_QString = QString::fromUtf8(kb.data, kb.len);
	self->setKeyBindings(kb_QString);
}

void QTermWidget_clear(QTermWidget* self) {
	self->clear();
}

void QTermWidget_toggleShowSearchBar(QTermWidget* self) {
	self->toggleShowSearchBar();
}

void QTermWidget_mouseEventEmiter(QTermWidget* self) {
	self->mouseEventEmiter();
}

struct miqt_string QTermWidget_tr2(const char* s, const char* c) {
	QString _ret = QTermWidget::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QTermWidget_tr3(const char* s, const char* c, int n) {
	QString _ret = QTermWidget::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QTermWidget_trUtf82(const char* s, const char* c) {
	QString _ret = QTermWidget::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QTermWidget_trUtf83(const char* s, const char* c, int n) {
	QString _ret = QTermWidget::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QTermWidget_selectedTextWithPreserveLineBreaks(QTermWidget* self, bool preserveLineBreaks) {
	QString _ret = self->selectedText(preserveLineBreaks);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QTermWidget_override_virtual_resizeEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__resizeEvent = slot;
	return true;
}

void QTermWidget_virtualbase_resizeEvent(void* self, QResizeEvent* param1) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::resizeEvent(param1);

}

bool QTermWidget_override_virtual_keyPressEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__keyPressEvent = slot;
	return true;
}

void QTermWidget_virtualbase_keyPressEvent(void* self, QKeyEvent* event) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::keyPressEvent(event);

}

bool QTermWidget_override_virtual_devType(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__devType = slot;
	return true;
}

int QTermWidget_virtualbase_devType(const void* self) {

	return ( (const MiqtVirtualQTermWidget*)(self) )->QTermWidget::devType();

}

bool QTermWidget_override_virtual_setVisible(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__setVisible = slot;
	return true;
}

void QTermWidget_virtualbase_setVisible(void* self, bool visible) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::setVisible(visible);

}

bool QTermWidget_override_virtual_sizeHint(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__sizeHint = slot;
	return true;
}

QSize* QTermWidget_virtualbase_sizeHint(const void* self) {

	return new QSize(( (const MiqtVirtualQTermWidget*)(self) )->QTermWidget::sizeHint());

}

bool QTermWidget_override_virtual_minimumSizeHint(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__minimumSizeHint = slot;
	return true;
}

QSize* QTermWidget_virtualbase_minimumSizeHint(const void* self) {

	return new QSize(( (const MiqtVirtualQTermWidget*)(self) )->QTermWidget::minimumSizeHint());

}

bool QTermWidget_override_virtual_heightForWidth(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__heightForWidth = slot;
	return true;
}

int QTermWidget_virtualbase_heightForWidth(const void* self, int param1) {

	return ( (const MiqtVirtualQTermWidget*)(self) )->QTermWidget::heightForWidth(static_cast<int>(param1));

}

bool QTermWidget_override_virtual_hasHeightForWidth(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__hasHeightForWidth = slot;
	return true;
}

bool QTermWidget_virtualbase_hasHeightForWidth(const void* self) {

	return ( (const MiqtVirtualQTermWidget*)(self) )->QTermWidget::hasHeightForWidth();

}

bool QTermWidget_override_virtual_paintEngine(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__paintEngine = slot;
	return true;
}

QPaintEngine* QTermWidget_virtualbase_paintEngine(const void* self) {

	return ( (const MiqtVirtualQTermWidget*)(self) )->QTermWidget::paintEngine();

}

bool QTermWidget_override_virtual_event(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__event = slot;
	return true;
}

bool QTermWidget_virtualbase_event(void* self, QEvent* event) {

	return ( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::event(event);

}

bool QTermWidget_override_virtual_mousePressEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__mousePressEvent = slot;
	return true;
}

void QTermWidget_virtualbase_mousePressEvent(void* self, QMouseEvent* event) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::mousePressEvent(event);

}

bool QTermWidget_override_virtual_mouseReleaseEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__mouseReleaseEvent = slot;
	return true;
}

void QTermWidget_virtualbase_mouseReleaseEvent(void* self, QMouseEvent* event) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::mouseReleaseEvent(event);

}

bool QTermWidget_override_virtual_mouseDoubleClickEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__mouseDoubleClickEvent = slot;
	return true;
}

void QTermWidget_virtualbase_mouseDoubleClickEvent(void* self, QMouseEvent* event) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::mouseDoubleClickEvent(event);

}

bool QTermWidget_override_virtual_mouseMoveEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__mouseMoveEvent = slot;
	return true;
}

void QTermWidget_virtualbase_mouseMoveEvent(void* self, QMouseEvent* event) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::mouseMoveEvent(event);

}

bool QTermWidget_override_virtual_wheelEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__wheelEvent = slot;
	return true;
}

void QTermWidget_virtualbase_wheelEvent(void* self, QWheelEvent* event) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::wheelEvent(event);

}

bool QTermWidget_override_virtual_keyReleaseEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__keyReleaseEvent = slot;
	return true;
}

void QTermWidget_virtualbase_keyReleaseEvent(void* self, QKeyEvent* event) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::keyReleaseEvent(event);

}

bool QTermWidget_override_virtual_focusInEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__focusInEvent = slot;
	return true;
}

void QTermWidget_virtualbase_focusInEvent(void* self, QFocusEvent* event) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::focusInEvent(event);

}

bool QTermWidget_override_virtual_focusOutEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__focusOutEvent = slot;
	return true;
}

void QTermWidget_virtualbase_focusOutEvent(void* self, QFocusEvent* event) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::focusOutEvent(event);

}

bool QTermWidget_override_virtual_enterEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__enterEvent = slot;
	return true;
}

void QTermWidget_virtualbase_enterEvent(void* self, QEvent* event) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::enterEvent(event);

}

bool QTermWidget_override_virtual_leaveEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__leaveEvent = slot;
	return true;
}

void QTermWidget_virtualbase_leaveEvent(void* self, QEvent* event) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::leaveEvent(event);

}

bool QTermWidget_override_virtual_paintEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__paintEvent = slot;
	return true;
}

void QTermWidget_virtualbase_paintEvent(void* self, QPaintEvent* event) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::paintEvent(event);

}

bool QTermWidget_override_virtual_moveEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__moveEvent = slot;
	return true;
}

void QTermWidget_virtualbase_moveEvent(void* self, QMoveEvent* event) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::moveEvent(event);

}

bool QTermWidget_override_virtual_closeEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__closeEvent = slot;
	return true;
}

void QTermWidget_virtualbase_closeEvent(void* self, QCloseEvent* event) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::closeEvent(event);

}

bool QTermWidget_override_virtual_contextMenuEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__contextMenuEvent = slot;
	return true;
}

void QTermWidget_virtualbase_contextMenuEvent(void* self, QContextMenuEvent* event) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::contextMenuEvent(event);

}

bool QTermWidget_override_virtual_tabletEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__tabletEvent = slot;
	return true;
}

void QTermWidget_virtualbase_tabletEvent(void* self, QTabletEvent* event) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::tabletEvent(event);

}

bool QTermWidget_override_virtual_actionEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__actionEvent = slot;
	return true;
}

void QTermWidget_virtualbase_actionEvent(void* self, QActionEvent* event) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::actionEvent(event);

}

bool QTermWidget_override_virtual_dragEnterEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__dragEnterEvent = slot;
	return true;
}

void QTermWidget_virtualbase_dragEnterEvent(void* self, QDragEnterEvent* event) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::dragEnterEvent(event);

}

bool QTermWidget_override_virtual_dragMoveEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__dragMoveEvent = slot;
	return true;
}

void QTermWidget_virtualbase_dragMoveEvent(void* self, QDragMoveEvent* event) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::dragMoveEvent(event);

}

bool QTermWidget_override_virtual_dragLeaveEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__dragLeaveEvent = slot;
	return true;
}

void QTermWidget_virtualbase_dragLeaveEvent(void* self, QDragLeaveEvent* event) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::dragLeaveEvent(event);

}

bool QTermWidget_override_virtual_dropEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__dropEvent = slot;
	return true;
}

void QTermWidget_virtualbase_dropEvent(void* self, QDropEvent* event) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::dropEvent(event);

}

bool QTermWidget_override_virtual_showEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__showEvent = slot;
	return true;
}

void QTermWidget_virtualbase_showEvent(void* self, QShowEvent* event) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::showEvent(event);

}

bool QTermWidget_override_virtual_hideEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__hideEvent = slot;
	return true;
}

void QTermWidget_virtualbase_hideEvent(void* self, QHideEvent* event) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::hideEvent(event);

}

bool QTermWidget_override_virtual_nativeEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__nativeEvent = slot;
	return true;
}

bool QTermWidget_virtualbase_nativeEvent(void* self, struct miqt_string eventType, void* message, long* result) {
	QByteArray eventType_QByteArray(eventType.data, eventType.len);

	return ( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::nativeEvent(eventType_QByteArray, message, static_cast<long*>(result));

}

bool QTermWidget_override_virtual_changeEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__changeEvent = slot;
	return true;
}

void QTermWidget_virtualbase_changeEvent(void* self, QEvent* param1) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::changeEvent(param1);

}

bool QTermWidget_override_virtual_metric(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__metric = slot;
	return true;
}

int QTermWidget_virtualbase_metric(const void* self, int param1) {

	return ( (const MiqtVirtualQTermWidget*)(self) )->QTermWidget::metric(static_cast<MiqtVirtualQTermWidget::PaintDeviceMetric>(param1));

}

bool QTermWidget_override_virtual_initPainter(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__initPainter = slot;
	return true;
}

void QTermWidget_virtualbase_initPainter(const void* self, QPainter* painter) {

	( (const MiqtVirtualQTermWidget*)(self) )->QTermWidget::initPainter(painter);

}

bool QTermWidget_override_virtual_redirected(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__redirected = slot;
	return true;
}

QPaintDevice* QTermWidget_virtualbase_redirected(const void* self, QPoint* offset) {

	return ( (const MiqtVirtualQTermWidget*)(self) )->QTermWidget::redirected(offset);

}

bool QTermWidget_override_virtual_sharedPainter(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__sharedPainter = slot;
	return true;
}

QPainter* QTermWidget_virtualbase_sharedPainter(const void* self) {

	return ( (const MiqtVirtualQTermWidget*)(self) )->QTermWidget::sharedPainter();

}

bool QTermWidget_override_virtual_inputMethodEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__inputMethodEvent = slot;
	return true;
}

void QTermWidget_virtualbase_inputMethodEvent(void* self, QInputMethodEvent* param1) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::inputMethodEvent(param1);

}

bool QTermWidget_override_virtual_inputMethodQuery(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__inputMethodQuery = slot;
	return true;
}

QVariant* QTermWidget_virtualbase_inputMethodQuery(const void* self, int param1) {

	return new QVariant(( (const MiqtVirtualQTermWidget*)(self) )->QTermWidget::inputMethodQuery(static_cast<Qt::InputMethodQuery>(param1)));

}

bool QTermWidget_override_virtual_focusNextPrevChild(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__focusNextPrevChild = slot;
	return true;
}

bool QTermWidget_virtualbase_focusNextPrevChild(void* self, bool next) {

	return ( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::focusNextPrevChild(next);

}

bool QTermWidget_override_virtual_eventFilter(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__eventFilter = slot;
	return true;
}

bool QTermWidget_virtualbase_eventFilter(void* self, QObject* watched, QEvent* event) {

	return ( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::eventFilter(watched, event);

}

bool QTermWidget_override_virtual_timerEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__timerEvent = slot;
	return true;
}

void QTermWidget_virtualbase_timerEvent(void* self, QTimerEvent* event) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::timerEvent(event);

}

bool QTermWidget_override_virtual_childEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__childEvent = slot;
	return true;
}

void QTermWidget_virtualbase_childEvent(void* self, QChildEvent* event) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::childEvent(event);

}

bool QTermWidget_override_virtual_customEvent(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__customEvent = slot;
	return true;
}

void QTermWidget_virtualbase_customEvent(void* self, QEvent* event) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::customEvent(event);

}

bool QTermWidget_override_virtual_connectNotify(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__connectNotify = slot;
	return true;
}

void QTermWidget_virtualbase_connectNotify(void* self, QMetaMethod* signal) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::connectNotify(*signal);

}

bool QTermWidget_override_virtual_disconnectNotify(void* self, intptr_t slot) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		return false;
	}
	
	self_cast->handle__disconnectNotify = slot;
	return true;
}

void QTermWidget_virtualbase_disconnectNotify(void* self, QMetaMethod* signal) {

	( (MiqtVirtualQTermWidget*)(self) )->QTermWidget::disconnectNotify(*signal);

}

void QTermWidget_protectedbase_sessionFinished(bool* _dynamic_cast_ok, void* self) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		*_dynamic_cast_ok = false;
		return ;
	}
	
	*_dynamic_cast_ok = true;
	
	self_cast->sessionFinished();

}

void QTermWidget_protectedbase_selectionChanged(bool* _dynamic_cast_ok, void* self, bool textSelected) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		*_dynamic_cast_ok = false;
		return ;
	}
	
	*_dynamic_cast_ok = true;
	
	self_cast->selectionChanged(textSelected);

}

void QTermWidget_protectedbase_updateMicroFocus(bool* _dynamic_cast_ok, void* self) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		*_dynamic_cast_ok = false;
		return ;
	}
	
	*_dynamic_cast_ok = true;
	
	self_cast->updateMicroFocus();

}

void QTermWidget_protectedbase_create(bool* _dynamic_cast_ok, void* self) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		*_dynamic_cast_ok = false;
		return ;
	}
	
	*_dynamic_cast_ok = true;
	
	self_cast->create();

}

void QTermWidget_protectedbase_destroy(bool* _dynamic_cast_ok, void* self) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		*_dynamic_cast_ok = false;
		return ;
	}
	
	*_dynamic_cast_ok = true;
	
	self_cast->destroy();

}

bool QTermWidget_protectedbase_focusNextChild(bool* _dynamic_cast_ok, void* self) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		*_dynamic_cast_ok = false;
		return false;
	}
	
	*_dynamic_cast_ok = true;
	
	return self_cast->focusNextChild();

}

bool QTermWidget_protectedbase_focusPreviousChild(bool* _dynamic_cast_ok, void* self) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		*_dynamic_cast_ok = false;
		return false;
	}
	
	*_dynamic_cast_ok = true;
	
	return self_cast->focusPreviousChild();

}

QObject* QTermWidget_protectedbase_sender(bool* _dynamic_cast_ok, const void* self) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		*_dynamic_cast_ok = false;
		return nullptr;
	}
	
	*_dynamic_cast_ok = true;
	
	return self_cast->sender();

}

int QTermWidget_protectedbase_senderSignalIndex(bool* _dynamic_cast_ok, const void* self) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		*_dynamic_cast_ok = false;
		return 0;
	}
	
	*_dynamic_cast_ok = true;
	
	return self_cast->senderSignalIndex();

}

int QTermWidget_protectedbase_receivers(bool* _dynamic_cast_ok, const void* self, const char* signal) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		*_dynamic_cast_ok = false;
		return 0;
	}
	
	*_dynamic_cast_ok = true;
	
	return self_cast->receivers(signal);

}

bool QTermWidget_protectedbase_isSignalConnected(bool* _dynamic_cast_ok, const void* self, QMetaMethod* signal) {
	MiqtVirtualQTermWidget* self_cast = dynamic_cast<MiqtVirtualQTermWidget*>( (QTermWidget*)(self) );
	if (self_cast == nullptr) {
		*_dynamic_cast_ok = false;
		return false;
	}
	
	*_dynamic_cast_ok = true;
	
	return self_cast->isSignalConnected(*signal);

}

void QTermWidget_delete(QTermWidget* self) {
	delete self;
}

