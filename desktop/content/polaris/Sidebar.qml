import QtQuick 2.15
import QtQuick.Controls 2.5
import QtQuick.Layouts 1.3

import quick 1.0

ColumnLayout {
    anchors.fill: parent
    spacing: 0

    property string currentPartition: "示例分区"
    property string currentLibrary: "示例笔记库"
    property bool showLibrarySelector: false

    Rectangle {
        Layout.preferredHeight: 32
        Layout.preferredWidth: parent.width
        Layout.alignment: Qt.AlignTop
        color: "#f8f8f8"

        RowLayout {
            anchors.fill: parent

            Rectangle {
                Layout.preferredWidth: 4
            }

            Text {
                Layout.alignment: Qt.AlignCenter | Qt.AlignLeft
                Layout.preferredWidth: 160
                text: currentLibrary
                font.pixelSize: 14
            }

            Rectangle {
                Layout.alignment: Qt.AlignCenter
                Layout.preferredHeight: 24
                Layout.preferredWidth: 24
                color: "transparent"
                Image {
                    anchors.fill: parent
                    source: "qrc:/qt/qml/quick/content/assets/material/symbols/web/keyboard_arrow_down/keyboard_arrow_down_48px.svg"
                    MouseArea {
                        anchors.fill: parent
                        onClicked: () => showLibrarySelector = !showLibrarySelector
                    }
                }
            }

            Rectangle {
                Layout.preferredWidth: 4
            }
        }
    }

    function selectLibrary(library) {
        currentLibrary = library
        showLibrarySelector = false
    }

    Rectangle {
        id: librarySelector
        Layout.preferredHeight: parent.height - 32
        Layout.preferredWidth: parent.width
        Layout.alignment: Qt.AlignTop
        visible: showLibrarySelector

        ColumnLayout {
            anchors.left: parent.left
            anchors.right: parent.right
            spacing: 0
            LibraryItem {
                library: "第一个笔记库"
                onCurrentLibraryChanged: library => selectLibrary(library)
            }
            LibraryItem {
                library: "另一个笔记库"
                onCurrentLibraryChanged: library => selectLibrary(library)
            }
            LibraryItem {
                library: "还有一个图片库"
                onCurrentLibraryChanged: library => selectLibrary(library)
            }
        }
    }

    Rectangle {
        id: partitionSelector
        Layout.preferredHeight: parent.height - 32
        Layout.preferredWidth: parent.width
        Layout.alignment: Qt.AlignTop
        visible: !showLibrarySelector

        // ColumnLayout {
        //     anchors.left: parent.left
        //     anchors.right: parent.right
        //     spacing: 0

        //     PartitionItem {
        //         partition: "第一个分区"
        //         onCurrentPartitionChanged: partition => currentPartition = partition
        //     }
        //     PartitionItem {
        //         partition: "第二个分区"
        //         onCurrentPartitionChanged: partition => currentPartition = partition
        //     }
        // }

        // Component {
        //     id: partitionDelegate
        //     delegate: PartitionItem {
        //         partition: 'name'
        //         onCurrentPartitionChanged: partition => currentPartition = partition
        //     }
        // }
        ListView {
            anchors.fill: parent
            anchors.left: parent.left
            anchors.right: parent.right
            boundsBehavior: Flickable.StopAtBounds
            model: PartitionModel {}
            delegate: PartitionItem {
                onCurrentPartitionChanged: partition => {
                                               console.log('partition',
                                                           partition)
                                               currentPartition = partition
                                           }
            }
        }
    }
}
