;;In the original input, replace [ with (, ] with ), and commas with spaces

;;;returns 1 if l1 is lower (correct order)
;;;returns 0 if equal
;;;returns -1 if l2 is lower
(defun listComp (l1 l2)
    (cond
    ((and (endp l1) (endp l2)) 0)
    ((endp l1) 1)
    ((endp l2) -1)
    ((and (numberp (first l1)) (numberp (first l2))) 
        (cond
        ((< (first l1) (first l2)) 1)
        ((> (first l1) (first l2)) -1)
        ((equal (first l1) (first l2)) (listComp (rest l1) (rest l2)))
        )
    )
    ((and (listp (first l1)) (listp (first l2))) 
        (cond
        ((equal (listComp (first l1) (first l2)) 0) (listComp (rest l1) (rest l2)))
        (T (listComp (first l1) (first l2)))
        ))
    ((and (listp (first l1)) (numberp (first l2)))
        (cond
        ((equal (listComp (first l1) (list (first l2))) 0) (listComp (rest l1) (rest l2)))
        (T (listComp (first l1) (list (first l2))))
        ))
    ((and (numberp (first l1)) (listp (first l2)))
        (cond
        ((equal (listComp (list (first l1)) (first l2)) 0) (listComp (rest l1) (rest l2)))
        (T (listComp (list (first l1)) (first l2)))
        ))
    (T (listComp (rest l1) (rest l2)))
    )
)

;PART 1
(setf score 0)
(with-open-file (stream "./13.txt")
 (dotimes (lines 150)
 (if (equal (listComp (read stream) (read stream)) 1) (setf score (+ score lines 1)) NIL)
 )
 (prin1 score)
)

;PART 2
;(defun twoComp (ls)
;    (cond
;    ((equal (listComp ls '((2))) 1) 2)
;    ((equal (listComp ls '((6))) 1) 1)
;    )
;)
;(setf scor1 0)
;(setf scor2 0)
;(with-open-file (stream "./13.txt")
; (dotimes (lines 300)
;    (setf comp (twoComp (read stream)))
;    (cond 
;    ((equal comp 2) (setf scor2 (+ scor2 1)))
;    ((equal comp 1) (setf scor1 (+ scor1 1)))
;    )
; )
; (prin1 (* (+ scor2 1) (+ scor2 scor1 2)))
;)
