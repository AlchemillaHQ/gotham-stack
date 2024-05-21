import DOMPurify from 'isomorphic-dompurify';
import { Notyf } from 'notyf';
import 'notyf/notyf.min.css';
import '../../css/custom-notyf.css'

window.htmx = require("htmx.org");
window.DOMPurify = DOMPurify;
window.notyf = new Notyf({
    duration: 1000,
    position: {
        x: 'center',
        y: 'bottom',
    },
    ripple: true,
    types: [
        {
            type: 'error',
            background: 'indianred',
            duration: 2500,
            dismissible: true
        },
        {
            type: 'success',
            background: 'seagreen',
            duration: 2500,
            dismissible: true
        },
        {
            type: 'info',
            background: 'royalblue',
            duration: 2500,
            dismissible: true
        }
    ]
});