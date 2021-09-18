    processor 6502

    include "vcs.h"
    include "macro.h"

    seg code
    org $f000               ; defines the origin of the ROM at $F000

START:
    ;CLEAN_START             ; Macro to safely clear the memory

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;; Set background luminosity to yello
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
    lda #$1E                ; LOad color into A ($1E is NTSC yellow)
    sta COLUBK              ; store A to backgroundColor Addres $09

    jmp START               ; Repeat from START

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;; Fill ROM to exactly 4KB
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
    org $FFFC               ; Defines origin to $FFFC
    .word START             ; Reset vector at $FFFC (where program starts)
    .word START             ; Interrupt vetor at $FFFE (unused in the VCS)

; EOF
